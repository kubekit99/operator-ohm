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

package openstackhelm

import (
	"context"
	"fmt"
	"time"

	"github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/helmv2"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	osh "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/apis/openstackhelm/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("controller_openstackhelm")

// Add creates a new OpenstackChart Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {

	return &HelmOperatorReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("openstackbackup-recorder"),
		managerFactory: helmv2.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
		watchDependentResources: true,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("openstackhelm-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource OpenstackChart
	err = c.Watch(&source.Kind{Type: &osh.OpenstackChart{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner OpenstackChart
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &osh.OpenstackChart{},
	})
	if err != nil {
		return err
	}

	// if options.WatchDependentResources {
	//	watchDependentResources(mgr, r, c)
	// }

	return nil
}

var _ reconcile.Reconciler = &HelmOperatorReconciler{}

// HelmOperatorReconciler reconciles custom resources as Helm releases.
type HelmOperatorReconciler struct {
	client          client.Client
	scheme          *runtime.Scheme
	recorder        record.EventRecorder
	gvk             schema.GroupVersionKind
	managerFactory  helmv2.ManagerFactory
	reconcilePeriod time.Duration
	//JEB releaseHook             ReleaseHookFunc
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
	instance := &osh.OpenstackChart{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)
	log := log.WithValues(
		"namespace", instance.GetNamespace(),
		"name", instance.GetName(),
	)
	log.V(1).Info("Reconciling")

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if apierrors.IsNotFound(err) {
		return reconcile.Result{}, nil
	}
	if err != nil {
		log.Error(err, "Failed to lookup resource")
		return reconcile.Result{}, err
	}

	manager := r.managerFactory.NewManager(instance)
	status := osh.StatusFor(instance)
	log = log.WithValues("release", manager.ReleaseName())

	deleted := instance.GetDeletionTimestamp() != nil
	pendingFinalizers := instance.GetFinalizers()
	if !deleted && !contains(pendingFinalizers, finalizer) {
		log.V(1).Info("Adding finalizer", "finalizer", finalizer)
		finalizers := append(pendingFinalizers, finalizer)
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	status.SetCondition(osh.OpenstackChartCondition{
		Type:   osh.ConditionInitialized,
		Status: osh.StatusTrue,
	})

	if err := manager.Sync(context.TODO()); err != nil {
		log.Error(err, "Failed to sync release")
		status.SetCondition(osh.OpenstackChartCondition{
			Type:    osh.ConditionIrreconcilable,
			Status:  osh.StatusTrue,
			Reason:  osh.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(osh.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizer) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledRelease, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != helmv2.ErrNotFound {
			log.Error(err, "Failed to uninstall release")
			status.SetCondition(osh.OpenstackChartCondition{
				Type:    osh.ConditionReleaseFailed,
				Status:  osh.StatusTrue,
				Reason:  osh.ReasonUninstallError,
				Message: err.Error(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(osh.ConditionReleaseFailed)

		if err == helmv2.ErrNotFound {
			log.Info("Release not found, removing finalizer")
		} else {
			log.Info("Uninstalled release")
			if log.Enabled() {
				fmt.Println(Diff(uninstalledRelease.GetManifest(), ""))
			}
			status.SetCondition(osh.OpenstackChartCondition{
				Type:   osh.ConditionDeployed,
				Status: osh.StatusFalse,
				Reason: osh.ReasonUninstallSuccessful,
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
			status.SetCondition(osh.OpenstackChartCondition{
				Type:           osh.ConditionReleaseFailed,
				Status:         osh.StatusTrue,
				Reason:         osh.ReasonInstallError,
				Message:        err.Error(),
				ReleaseName:    installedRelease.GetName(),
				ReleaseVersion: installedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(osh.ConditionReleaseFailed)

		// if r.releaseHook != nil {
		// 	if err := r.releaseHook(installedRelease); err != nil {
		// 		log.Error(err, "Failed to run release hook")
		// 		return reconcile.Result{}, err
		// 	}
		// }

		log.Info("Installed release")
		if log.Enabled() {
			fmt.Println(Diff("", installedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", installedRelease.GetConfig())
		status.SetCondition(osh.OpenstackChartCondition{
			Type:           osh.ConditionDeployed,
			Status:         osh.StatusTrue,
			Reason:         osh.ReasonInstallSuccessful,
			Message:        installedRelease.GetInfo().GetStatus().GetNotes(),
			ReleaseName:    installedRelease.GetName(),
			ReleaseVersion: installedRelease.GetVersion(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	if manager.IsUpdateRequired() {
		previousRelease, updatedRelease, err := manager.UpdateRelease(context.TODO())
		if err != nil {
			log.Error(err, "Failed to update release")
			status.SetCondition(osh.OpenstackChartCondition{
				Type:           osh.ConditionReleaseFailed,
				Status:         osh.StatusTrue,
				Reason:         osh.ReasonUpdateError,
				Message:        err.Error(),
				ReleaseName:    updatedRelease.GetName(),
				ReleaseVersion: updatedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(osh.ConditionReleaseFailed)

		// if r.releaseHook != nil {
		// 	if err := r.releaseHook(updatedRelease); err != nil {
		// 		log.Error(err, "Failed to run release hook")
		// 		return reconcile.Result{}, err
		// 	}
		// }

		log.Info("Updated release")
		if log.Enabled() {
			fmt.Println(Diff(previousRelease.GetManifest(), updatedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", updatedRelease.GetConfig())
		status.SetCondition(osh.OpenstackChartCondition{
			Type:           osh.ConditionDeployed,
			Status:         osh.StatusTrue,
			Reason:         osh.ReasonUpdateSuccessful,
			Message:        updatedRelease.GetInfo().GetStatus().GetNotes(),
			ReleaseName:    updatedRelease.GetName(),
			ReleaseVersion: updatedRelease.GetVersion(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	// expectedRelease, err := manager.ReconcileRelease(context.TODO())
	_, err = manager.ReconcileRelease(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile release")
		status.SetCondition(osh.OpenstackChartCondition{
			Type:    osh.ConditionIrreconcilable,
			Status:  osh.StatusTrue,
			Reason:  osh.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(osh.ConditionIrreconcilable)

	// if r.releaseHook != nil {
	// 	if err := r.releaseHook(expectedRelease); err != nil {
	// 		log.Error(err, "Failed to run release hook")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	log.Info("Reconciled release")
	err = r.updateResourceStatus(instance, status)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

func (r HelmOperatorReconciler) updateResource(o *osh.OpenstackChart) error {
	return r.client.Update(context.TODO(), o)
}

func (r HelmOperatorReconciler) updateResourceStatus(instance *osh.OpenstackChart, status *osh.OpenstackChartStatus) error {
	reqLogger := log.WithValues("OpenstackChart.Namespace", instance.Namespace, "OpenstackChart.Name", instance.Name)

	// JEB: This is already a reference to the object
	// instance.Status = status

	// JEB: Be sure to have update status subresources in the CRD.yaml
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status")
	}

	return err
}
