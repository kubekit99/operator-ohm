// Copyright 2019 The Armada Authors
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

package armada

import (
	"context"
	"fmt"
	"time"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	armadamgr "github.com/kubekit99/operator-ohm/armada-operator/pkg/armada"
	armadaif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a new ArmadaManifest Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddArmadaManifestController(mgr manager.Manager) error {

	r := &ArmadaManifestReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("amf-recorder"),
		managerFactory: armadamgr.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
	}

	// Create a new controller
	c, err := controller.New("amf-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArmadaManifest
	err = c.Watch(&source.Kind{Type: &av1.ArmadaManifest{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ArmadaManifest
	owner := av1.NewArmadaManifestVersionKind("", "")
	r.depResourceWatchUpdater = armadaif.BuildDependentResourceWatchUpdater(mgr, owner, c)

	return nil
}

var _ reconcile.Reconciler = &ArmadaManifestReconciler{}

// ArmadaManifestReconciler reconciles a ArmadaManifest object
type ArmadaManifestReconciler struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          armadaif.ArmadaManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater armadaif.DependentResourceWatchUpdater
}

const (
	finalizerArmadaManifest = "uninstall-pod"
)

// Reconcile reads that state of the cluster for a ArmadaManifest object and makes changes based on the state read
// and what is in the ArmadaManifest.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ArmadaManifestReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &av1.ArmadaManifest{}
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

	// AdminState POC begin
	// We will have to enhance the placement of this test to account
	// for kubectl apply where more than just the AdminState is changed
	if disabled := r.isReconcileDisabled(instance); disabled {
		return reconcile.Result{}, nil
	}
	// AdminState POC end

	manager := r.managerFactory.NewArmadaManifestManager(instance)
	spec := instance.Spec
	status := &instance.Status

	log = log.WithValues("resource", manager.ResourceName())

	deleted := instance.GetDeletionTimestamp() != nil
	pendingFinalizers := instance.GetFinalizers()
	if !deleted && !contains(pendingFinalizers, finalizerArmadaManifest) {
		finalizers := append(pendingFinalizers, finalizerArmadaManifest)
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	hrc := av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	}
	status.SetCondition(hrc)
	status.ComputeActualState(&hrc, spec.TargetState)

	if err := manager.Sync(context.TODO()); err != nil {
		hrc := av1.HelmResourceCondition{
			Type:    av1.ConditionIrreconcilable,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonReconcileError,
			Message: err.Error(),
		}
		status.SetCondition(hrc)
		status.ComputeActualState(&hrc, spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(av1.ConditionIrreconcilable)

	if deleted {
		if !contains(pendingFinalizers, finalizerArmadaManifest) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledResource, err := manager.UninstallResource(context.TODO())
		if err != nil && err != armadaif.ErrNotFound {
			hrc := av1.HelmResourceCondition{
				Type:         av1.ConditionFailed,
				Status:       av1.ConditionStatusTrue,
				Reason:       av1.ReasonUninstallError,
				Message:      err.Error(),
				ResourceName: uninstalledResource.GetName(),
			}
			status.SetCondition(hrc)
			status.ComputeActualState(&hrc, spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if err == armadaif.ErrNotFound {
			log.Info("Resource not found, removing finalizer")
		} else {
			hrc := av1.HelmResourceCondition{
				Type:   av1.ConditionDeployed,
				Status: av1.ConditionStatusFalse,
				Reason: av1.ReasonUninstallSuccessful,
			}
			status.SetCondition(hrc)
			status.ComputeActualState(&hrc, spec.TargetState)
			r.logAndRecordSuccess(instance, &hrc)
		}
		if err := r.updateResourceStatus(instance, status); err != nil {
			return reconcile.Result{}, err
		}

		finalizers := []string{}
		for _, pendingFinalizer := range pendingFinalizers {
			if pendingFinalizer != finalizerArmadaManifest {
				finalizers = append(finalizers, pendingFinalizer)
			}
		}
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	if !manager.IsInstalled() {
		installedResource, err := manager.InstallResource(context.TODO())
		if err != nil {
			hrc := av1.HelmResourceCondition{
				Type:    av1.ConditionFailed,
				Status:  av1.ConditionStatusTrue,
				Reason:  av1.ReasonInstallError,
				Message: err.Error(),
			}
			status.SetCondition(hrc)
			status.ComputeActualState(&hrc, spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if r.depResourceWatchUpdater != nil {
			if err := r.depResourceWatchUpdater(instance.GetDependentResources()); err != nil {
				log.Error(err, "Failed to run update resource dependent resources")
				return reconcile.Result{}, err
			}
		}

		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionDeployed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonInstallSuccessful,
			Message:      "",
			ResourceName: installedResource.GetName(),
		}
		status.SetCondition(hrc)
		status.ComputeActualState(&hrc, spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)

		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	if manager.IsUpdateRequired() {
		previousResource, updatedResource, err := manager.UpdateResource(context.TODO())
		if previousResource != nil && updatedResource != nil {
			log.Info(previousResource.GetName(), updatedResource.GetName())
		}
		if err != nil {
			hrc := av1.HelmResourceCondition{
				Type:         av1.ConditionFailed,
				Status:       av1.ConditionStatusTrue,
				Reason:       av1.ReasonUpdateError,
				Message:      err.Error(),
				ResourceName: updatedResource.GetName(),
			}
			status.SetCondition(hrc)
			status.ComputeActualState(&hrc, spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		if r.depResourceWatchUpdater != nil {
			if err := r.depResourceWatchUpdater(instance.GetDependentResources()); err != nil {
				log.Error(err, "Failed to run update resource dependent resources")
				return reconcile.Result{}, err
			}
		}

		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionDeployed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUpdateSuccessful,
			Message:      "HardcodedMessage",
			ResourceName: updatedResource.GetName(),
		}
		status.SetCondition(hrc)
		status.ComputeActualState(&hrc, spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)

		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	expectedResource, err := manager.ReconcileResource(context.TODO())
	if err != nil {
		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionIrreconcilable,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonReconcileError,
			Message:      err.Error(),
			ResourceName: expectedResource.GetName(),
		}
		status.SetCondition(hrc)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance, status)
		return reconcile.Result{}, err
	}
	status.RemoveCondition(av1.ConditionIrreconcilable)

	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(instance.GetDependentResources()); err != nil {
			log.Error(err, "Failed to run update resource dependent resources")
			return reconcile.Result{}, err
		}
	}

	log.Info("Reconciled resource")
	err = r.updateResourceStatus(instance, status)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// Add a success event to the recorder
func (r ArmadaManifestReconciler) logAndRecordFailure(instance *av1.ArmadaManifest, hrc *av1.HelmResourceCondition, err error) {
	log.Error(err, fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

func (r ArmadaManifestReconciler) logAndRecordSuccess(instance *av1.ArmadaManifest, hrc *av1.HelmResourceCondition) {
	log.Info(fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// Update the Resource object
func (r ArmadaManifestReconciler) updateResource(o *av1.ArmadaManifest) error {
	return r.client.Update(context.TODO(), o)
}

// Update the Status field in the CRD
func (r ArmadaManifestReconciler) updateResourceStatus(instance *av1.ArmadaManifest, status *av1.ArmadaManifestStatus) error {
	reqLogger := log.WithValues("ArmadaManifest.Namespace", instance.Namespace, "ArmadaManifest.Name", instance.Name)

	helper := av1.HelmResourceConditionListHelper{Items: status.Conditions}
	status.Conditions = helper.InitIfEmpty()

	// JEB: Be sure to have update status subresources in the CRD.yaml
	// JEB: Look for kubebuilder subresources in the _types.go
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status. Ignoring")
		err = nil
	}

	return err
}

// isReconcileDisabled
func (r ArmadaManifestReconciler) isReconcileDisabled(instance *av1.ArmadaManifest) bool {
	// JEB: Not sure if we need to add this new ConditionEnabled
	// or we can just used the ConditionInitialized
	if instance.IsDisabled() {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusFalse,
			Reason: "Manifest is disabled",
		}
		r.logAndRecordSuccess(instance, &hrc)
		instance.Status.SetCondition(hrc)
		_ = r.updateResourceStatus(instance, &instance.Status)
		return true
	} else {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusTrue,
			Reason: "Manifest is enabled",
		}
		r.logAndRecordSuccess(instance, &hrc)
		instance.Status.SetCondition(hrc)
		return false
	}
}
