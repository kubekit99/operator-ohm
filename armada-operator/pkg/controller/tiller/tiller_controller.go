// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tiller

import (
	"context"
	"fmt"
	"time"

	helmmgr "github.com/kubekit99/operator-ohm/armada-operator/pkg/helm"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("tiller")

// Add creates a new HelmRelease Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr)
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager) error {
	r := &HelmOperatorReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("tiller-recorder"),
		managerFactory: helmmgr.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
		watchDependentResources: true,
	}

	// Create a new controller
	c, err := controller.New("tiller-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource HelmRelease
	err = c.Watch(&source.Kind{Type: &av1.HelmRelease{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner HelmRelease
	owner := av1.NewHelmReleaseVersionKind()
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    owner,
	})
	if err != nil {
		return err
	}

	r.releaseWatchUpdater = helmif.BuildReleaseDependantResourcesWatchUpdater(mgr, owner, c)

	return nil
}

var _ reconcile.Reconciler = &HelmOperatorReconciler{}

// HelmOperatorReconciler reconciles custom resources as Helm releases.
type HelmOperatorReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          helmif.HelmManagerFactory
	reconcilePeriod         time.Duration
	releaseWatchUpdater     helmif.ReleaseWatchUpdater
	watchDependentResources bool
}

const (
	finalizer = "uninstall-helm-release"
)

// Reconcile reconciles the requested resource by installing, updating, or
// uninstalling a Helm release based on the resource's current state. If no
// release changes are necessary, Reconcile will create or patch the underlying
// resources to match the expected release manifest.
func (r HelmOperatorReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &av1.HelmRelease{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)
	log := log.WithValues(
		"namespace", instance.GetNamespace(),
		"name", instance.GetName(),
	)

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if apierrors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}
	if err != nil {
		log.Error(err, "Failed to lookup resource")
		return reconcile.Result{}, err
	}

	manager := r.managerFactory.NewTillerManager(instance)
	spec := instance.Spec
	status := &instance.Status

	log = log.WithValues("release", manager.ReleaseName())

	deleted := instance.GetDeletionTimestamp() != nil
	pendingFinalizers := instance.GetFinalizers()
	if !deleted && !contains(pendingFinalizers, finalizer) {
		finalizers := append(pendingFinalizers, finalizer)
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	status.SetCondition(av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	})

	if err := manager.Sync(context.TODO()); err != nil {
		log.Error(err, "Failed to sync release")
		status.SetCondition(av1.HelmResourceCondition{
			Type:    av1.ConditionIrreconcilable,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(av1.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizer) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledRelease, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != helmif.ErrNotFound {
			log.Error(err, "Failed to uninstall release")
			status.SetCondition(av1.HelmResourceCondition{
				Type:    av1.ConditionFailed,
				Status:  av1.ConditionStatusTrue,
				Reason:  av1.ReasonUninstallError,
				Message: err.Error(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if err == helmif.ErrNotFound {
			log.Info("Release not found, removing finalizer")
		} else {
			r.recorder.Event(instance, v1.EventTypeWarning, "DeletionFailure", fmt.Sprintf("Uninstalled Release %s", uninstalledRelease.GetName()))
			log.Info("Uninstalled release", "releaseName", uninstalledRelease.GetName(), "releaseVersion", uninstalledRelease.GetVersion())
			status.SetCondition(av1.HelmResourceCondition{
				Type:   av1.ConditionDeployed,
				Status: av1.ConditionStatusFalse,
				Reason: av1.ReasonUninstallSuccessful,
			})
		}
		if err := r.updateResourceStatus(instance, status); err != nil {
			return reconcile.Result{}, err
		}

		finalizers := []string{}
		for _, pendingFinalizer := range pendingFinalizers {
			if pendingFinalizer != finalizer {
				finalizers = append(finalizers, pendingFinalizer)
			}
		}
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	if !manager.IsInstalled() {
		installedRelease, err := manager.InstallRelease(context.TODO())
		if err != nil {
			log.Error(err, "Failed to install release")
			r.recorder.Event(instance, v1.EventTypeWarning, "InstallationFailure", fmt.Sprintf("Installed Release %s", installedRelease.GetName()))
			status.SetCondition(av1.HelmResourceCondition{
				Type:            av1.ConditionFailed,
				Status:          av1.ConditionStatusTrue,
				Reason:          av1.ReasonInstallError,
				Message:         err.Error(),
				ResourceName:    installedRelease.GetName(),
				ResourceVersion: installedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if spec.WatchHelmDependentResources && r.releaseWatchUpdater != nil {
			if err := r.releaseWatchUpdater(installedRelease); err != nil {
				log.Error(err, "Failed to run update release dependant resources")
				return reconcile.Result{}, err
			}
		}

		log.Info("Installed release", "releaseName", installedRelease.GetName(), "releaseVersion", installedRelease.GetVersion())
		r.recorder.Event(instance, v1.EventTypeNormal, "Installed", fmt.Sprintf("Installed Release %s", installedRelease.GetName()))
		status.SetCondition(av1.HelmResourceCondition{
			Type:            av1.ConditionDeployed,
			Status:          av1.ConditionStatusTrue,
			Reason:          av1.ReasonInstallSuccessful,
			Message:         installedRelease.GetInfo().GetStatus().GetNotes(),
			ResourceName:    installedRelease.GetName(),
			ResourceVersion: installedRelease.GetVersion(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	if manager.IsUpdateRequired() {
		previousRelease, updatedRelease, err := manager.UpdateRelease(context.TODO())
		if err != nil {
			log.Error(err, "Failed to update release")
			r.recorder.Event(instance, v1.EventTypeWarning, "UpdateFailure", fmt.Sprintf("Updated Release %s", updatedRelease.GetName()))
			status.SetCondition(av1.HelmResourceCondition{
				Type:            av1.ConditionFailed,
				Status:          av1.ConditionStatusTrue,
				Reason:          av1.ReasonUpdateError,
				Message:         err.Error(),
				ResourceName:    updatedRelease.GetName(),
				ResourceVersion: updatedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if spec.WatchHelmDependentResources && r.releaseWatchUpdater != nil {
			if err := r.releaseWatchUpdater(updatedRelease); err != nil {
				log.Error(err, "Failed to run update release dependant resources")
				return reconcile.Result{}, err
			}
		}

		log.Info("Updated release", "releaseName", updatedRelease.GetName(), "releaseVersion", updatedRelease.GetVersion())
		r.recorder.Event(instance, v1.EventTypeNormal, "Updated", fmt.Sprintf("Updated Release %s", updatedRelease.GetName()))
		if log.Enabled() {
			fmt.Println(Diff(previousRelease.GetManifest(), updatedRelease.GetManifest()))
		}
		status.SetCondition(av1.HelmResourceCondition{
			Type:            av1.ConditionDeployed,
			Status:          av1.ConditionStatusTrue,
			Reason:          av1.ReasonUpdateSuccessful,
			Message:         updatedRelease.GetInfo().GetStatus().GetNotes(),
			ResourceName:    updatedRelease.GetName(),
			ResourceVersion: updatedRelease.GetVersion(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	expectedRelease, err := manager.ReconcileRelease(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile release")
		status.SetCondition(av1.HelmResourceCondition{
			Type:    av1.ConditionIrreconcilable,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(av1.ConditionIrreconcilable)

	if spec.WatchHelmDependentResources && r.releaseWatchUpdater != nil {
		if err := r.releaseWatchUpdater(expectedRelease); err != nil {
			log.Error(err, "Failed to run update release dependant resources")
			return reconcile.Result{}, err
		}
	}

	log.Info("Reconciled release")
	err = r.updateResourceStatus(instance, status)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// Update the whole CRD
func (r HelmOperatorReconciler) updateResource(o *av1.HelmRelease) error {
	return r.client.Update(context.TODO(), o)
}

// Update the Status field in the CRD
func (r HelmOperatorReconciler) updateResourceStatus(instance *av1.HelmRelease, status *av1.HelmReleaseStatus) error {
	reqLogger := log.WithValues("HelmRelease.Namespace", instance.Namespace, "HelmRelease.Name", instance.Name)

	helper := av1.HelmResourceConditionListHelper{Items: status.Conditions}
	status.Conditions = helper.InitIfEmpty()

	if log.Enabled() {
		fmt.Println(helper.PrettyPrint())
	}

	// JEB: Be sure to have update status subresources in the CRD.yaml
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status. Ignoring")
	}

	return nil
}
