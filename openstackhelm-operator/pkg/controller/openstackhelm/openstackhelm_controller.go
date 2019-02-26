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

	argoif "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/argoif"
	helmif "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/helmif"
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

	oshv1 "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/apis/openstackhelm/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("controller_openstackhelm")

// Add creates a new OpenstackChart Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr)
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager) error {
	r := &HelmOperatorReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("openstackbackup-recorder"),
		managerFactory: helmv2.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
		watchDependentResources: true,
	}

	// Create a new controller
	c, err := controller.New("openstackhelm-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource OpenstackChart
	err = c.Watch(&source.Kind{Type: &oshv1.OpenstackChart{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner OpenstackChart
	owner := oshv1.NewOpenstackChartVersionKind()
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    owner,
	})
	if err != nil {
		return err
	}

	r.releaseWatchUpdater = helmif.BuildReleaseDependantResourcesWatchUpdater(mgr, owner, c)

	// Watch for changes to secondary resource Workflows and requeue the owner OpenstackDeployment
	// JEB: When a Workflow owned by OpenstackChart is deleted, the Reconcile method is invoked.
	o := argoif.NewWorkflowGroupVersionKind()
	err = c.Watch(&source.Kind{Type: o},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    owner,
		},
		// predicate.GenerationChangedPredicate{}
	)
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &HelmOperatorReconciler{}

// HelmOperatorReconciler reconciles custom resources as Helm releases.
type HelmOperatorReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	gvk                     schema.GroupVersionKind
	managerFactory          helmif.ManagerFactory
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
	instance := &oshv1.OpenstackChart{}
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
	spec := instance.Spec
	status := oshv1.StatusFor(instance)
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

	status.SetCondition(oshv1.OpenstackChartCondition{
		Type:   oshv1.ConditionInitialized,
		Status: oshv1.StatusTrue,
	})

	if err := manager.Sync(context.TODO()); err != nil {
		log.Error(err, "Failed to sync release")
		status.SetCondition(oshv1.OpenstackChartCondition{
			Type:    oshv1.ConditionIrreconcilable,
			Status:  oshv1.StatusTrue,
			Reason:  oshv1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(oshv1.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizer) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledRelease, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != helmif.ErrNotFound {
			log.Error(err, "Failed to uninstall release")
			status.SetCondition(oshv1.OpenstackChartCondition{
				Type:    oshv1.ConditionReleaseFailed,
				Status:  oshv1.StatusTrue,
				Reason:  oshv1.ReasonUninstallError,
				Message: err.Error(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(oshv1.ConditionReleaseFailed)

		if err == helmif.ErrNotFound {
			log.Info("Release not found, removing finalizer")
		} else {
			log.Info("Uninstalled release")
			if log.Enabled() {
				fmt.Println(Diff(uninstalledRelease.GetManifest(), ""))
			}
			status.SetCondition(oshv1.OpenstackChartCondition{
				Type:   oshv1.ConditionDeployed,
				Status: oshv1.StatusFalse,
				Reason: oshv1.ReasonUninstallSuccessful,
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
			status.SetCondition(oshv1.OpenstackChartCondition{
				Type:           oshv1.ConditionReleaseFailed,
				Status:         oshv1.StatusTrue,
				Reason:         oshv1.ReasonInstallError,
				Message:        err.Error(),
				ReleaseName:    installedRelease.GetName(),
				ReleaseVersion: installedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(oshv1.ConditionReleaseFailed)

		if spec.WatchHelmDependentResources && r.releaseWatchUpdater != nil {
			if err := r.releaseWatchUpdater(installedRelease); err != nil {
				log.Error(err, "Failed to run update release dependant resources")
				return reconcile.Result{}, err
			}
		}

		log.Info("Installed release")
		if log.Enabled() {
			fmt.Println(Diff("", installedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", installedRelease.GetConfig())
		status.SetCondition(oshv1.OpenstackChartCondition{
			Type:           oshv1.ConditionDeployed,
			Status:         oshv1.StatusTrue,
			Reason:         oshv1.ReasonInstallSuccessful,
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
			status.SetCondition(oshv1.OpenstackChartCondition{
				Type:           oshv1.ConditionReleaseFailed,
				Status:         oshv1.StatusTrue,
				Reason:         oshv1.ReasonUpdateError,
				Message:        err.Error(),
				ReleaseName:    updatedRelease.GetName(),
				ReleaseVersion: updatedRelease.GetVersion(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(oshv1.ConditionReleaseFailed)

		if spec.WatchHelmDependentResources && r.releaseWatchUpdater != nil {
			if err := r.releaseWatchUpdater(updatedRelease); err != nil {
				log.Error(err, "Failed to run update release dependant resources")
				return reconcile.Result{}, err
			}
		}

		log.Info("Updated release")
		if log.Enabled() {
			fmt.Println(Diff(previousRelease.GetManifest(), updatedRelease.GetManifest()))
		}
		log.V(1).Info("Config values", "values", updatedRelease.GetConfig())
		status.SetCondition(oshv1.OpenstackChartCondition{
			Type:           oshv1.ConditionDeployed,
			Status:         oshv1.StatusTrue,
			Reason:         oshv1.ReasonUpdateSuccessful,
			Message:        updatedRelease.GetInfo().GetStatus().GetNotes(),
			ReleaseName:    updatedRelease.GetName(),
			ReleaseVersion: updatedRelease.GetVersion(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	expectedRelease, err := manager.ReconcileRelease(context.TODO())
	_, err = manager.ReconcileRelease(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile release")
		status.SetCondition(oshv1.OpenstackChartCondition{
			Type:    oshv1.ConditionIrreconcilable,
			Status:  oshv1.StatusTrue,
			Reason:  oshv1.ReasonReconcileError,
			Message: err.Error(),
		})
		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(oshv1.ConditionIrreconcilable)

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

func (r HelmOperatorReconciler) updateResource(o *oshv1.OpenstackChart) error {
	return r.client.Update(context.TODO(), o)
}

func (r HelmOperatorReconciler) updateResourceStatus(instance *oshv1.OpenstackChart, status *oshv1.OpenstackChartStatus) error {
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
