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

// AddArmadaChartGroupController creates a new ArmadaChartGroup Controller and
// adds it to the Manager. The Manager will set fields on the Controller and
// Start it when the Manager is Started.
func AddArmadaChartGroupController(mgr manager.Manager) error {

	r := &ChartGroupReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("acg-recorder"),
		managerFactory: armadamgr.NewManagerFactory(mgr),
	}

	// Create a new controller
	c, err := controller.New("acg-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArmadaChartGroup
	err = c.Watch(&source.Kind{Type: &av1.ArmadaChartGroup{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ArmadaChartGroup
	owner := av1.NewArmadaChartGroupVersionKind("", "")
	r.depResourceWatchUpdater = armadaif.BuildDependantResourceWatchUpdater(mgr, owner, c)

	return nil
}

var _ reconcile.Reconciler = &ChartGroupReconciler{}

// ChartGroupReconciler reconciles a ArmadaChartGroup object
type ChartGroupReconciler struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          armadaif.ArmadaManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater armadaif.DependantResourceWatchUpdater
}

const (
	finalizerArmadaChartGroup = "uninstall-pod"
)

// Reconcile reads that state of the cluster for a ArmadaChartGroup object and
// makes changes based on the state read and what is in the ArmadaChartGroup.Spec
//
// Note: The Controller will requeue the Request to be processed again if the
// returned error is non-nil or Result.Requeue is true, otherwise upon
// completion it will remove the work from the queue.
func (r *ChartGroupReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instance := &av1.ArmadaChartGroup{}
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

	mgr := r.managerFactory.NewArmadaChartGroupManager(instance)

	log = log.WithValues("resource", mgr.ResourceName())

	var shouldRequeue bool
	if shouldRequeue, err = r.updateFinalizers(instance); shouldRequeue {
		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	hrc := av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	}
	instance.Status.SetCondition(hrc)
	instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)

	if err := r.ensureSynced(mgr, instance); err != nil {
		return reconcile.Result{}, err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)

	switch {
	case instance.IsDeleted():
		if shouldRequeue, err = r.deleteArmadaChartGroup(mgr, instance); shouldRequeue {
			// Need to requeue because finalizer update does not change metadata.generation
			return reconcile.Result{Requeue: true}, err
		}
		return reconcile.Result{}, err
	case !mgr.IsInstalled():
		if shouldRequeue, err = r.installArmadaChartGroup(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	case mgr.IsUpdateRequired():
		if shouldRequeue, err = r.updateArmadaChartGroup(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	}

	if err := r.reconcileArmadaChartGroup(mgr, instance); err != nil {
		return reconcile.Result{}, err
	}

	log.Info("Reconciled resource")
	err = r.updateResourceStatus(instance)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// logAndRecordFailure adds a failure event to the recorder
func (r ChartGroupReconciler) logAndRecordFailure(instance *av1.ArmadaChartGroup, hrc *av1.HelmResourceCondition, err error) {
	log.Error(err, fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

// logAndRecordSuccess adds a success event to the recorder
func (r ChartGroupReconciler) logAndRecordSuccess(instance *av1.ArmadaChartGroup, hrc *av1.HelmResourceCondition) {
	log.Info(fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// updateResource updates the Resource object in the cluster
func (r ChartGroupReconciler) updateResource(o *av1.ArmadaChartGroup) error {
	return r.client.Update(context.TODO(), o)
}

// updateResourceStatus updates the the Status field of the Resource object in the cluster
func (r ChartGroupReconciler) updateResourceStatus(instance *av1.ArmadaChartGroup) error {
	reqLogger := log.WithValues("ArmadaChartGroup.Namespace", instance.Namespace, "ArmadaChartGroup.Name", instance.Name)

	helper := av1.HelmResourceConditionListHelper{Items: instance.Status.Conditions}
	instance.Status.Conditions = helper.InitIfEmpty()

	// JEB: Be sure to have update status subresources in the CRD.yaml
	// JEB: Look for kubebuilder subresources in the _types.go
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status. Ignoring")
		err = nil
	}

	return err
}

// ensureSynced checks that the ArmadaManager is in sync with the cluster
func (r ChartGroupReconciler) ensureSynced(mgr armadaif.ArmadaManager, instance *av1.ArmadaChartGroup) error {
	if err := mgr.Sync(context.TODO()); err != nil {
		hrc := av1.HelmResourceCondition{
			Type:    av1.ConditionIrreconcilable,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonReconcileError,
			Message: err.Error(),
		}
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return err
	}
	return nil
}

// updateFinalizers asserts that the finalizers match what is expected based on
// whether the instance is currently being deleted or not. It returns true if
// the finalizers were changed, false otherwise
func (r ChartGroupReconciler) updateFinalizers(instance *av1.ArmadaChartGroup) (bool, error) {
	pendingFinalizers := instance.GetFinalizers()
	if !instance.IsDeleted() && !contains(pendingFinalizers, finalizerArmadaChartGroup) {
		finalizers := append(pendingFinalizers, finalizerArmadaChartGroup)
		instance.SetFinalizers(finalizers)
		err := r.updateResource(instance)

		return true, err
	}
	return false, nil
}

// deleteArmadaChartGroup deletes an instance of an ArmadaChartGroup. It returns true if the reconciler should be re-enqueueed
func (r ChartGroupReconciler) deleteArmadaChartGroup(mgr armadaif.ArmadaManager, instance *av1.ArmadaChartGroup) (bool, error) {
	pendingFinalizers := instance.GetFinalizers()
	if !contains(pendingFinalizers, finalizerArmadaChartGroup) {
		log.Info("Resource is terminated, skipping reconciliation")
		return false, nil
	}

	uninstalledResource, err := mgr.UninstallResource(context.TODO())
	if err != nil && err != armadaif.ErrNotFound {
		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionFailed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUninstallError,
			Message:      err.Error(),
			ResourceName: uninstalledResource.GetName(),
		}
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err == armadaif.ErrNotFound {
		log.Info("Resource not found, removing finalizer")
	} else {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionDeployed,
			Status: av1.ConditionStatusFalse,
			Reason: av1.ReasonUninstallSuccessful,
		}
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)
	}
	if err := r.updateResourceStatus(instance); err != nil {
		return false, err
	}

	finalizers := []string{}
	for _, pendingFinalizer := range pendingFinalizers {
		if pendingFinalizer != finalizerArmadaChartGroup {
			finalizers = append(finalizers, pendingFinalizer)
		}
	}
	instance.SetFinalizers(finalizers)
	err = r.updateResource(instance)

	// Need to requeue because finalizer update does not change metadata.generation
	return true, err
}

// installArmadaChartGroup attempts to install instance. It returns true if the reconciler should be re-enqueueed
func (r ChartGroupReconciler) installArmadaChartGroup(mgr armadaif.ArmadaManager, instance *av1.ArmadaChartGroup) (bool, error) {
	installedResource, err := mgr.InstallResource(context.TODO())
	if err != nil {
		hrc := av1.HelmResourceCondition{
			Type:    av1.ConditionFailed,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonInstallError,
			Message: err.Error(),
		}
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(instance.GetDependantResources()); err != nil {
			log.Error(err, "Failed to run update resource dependant resources")
			return false, err
		}
	}

	hrc := av1.HelmResourceCondition{
		Type:         av1.ConditionDeployed,
		Status:       av1.ConditionStatusTrue,
		Reason:       av1.ReasonInstallSuccessful,
		Message:      "",
		ResourceName: installedResource.GetName(),
	}
	instance.Status.SetCondition(hrc)
	instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// updateArmadaChartGroup attempts to update instance. It returns true if the reconciler should be re-enqueueed
func (r ChartGroupReconciler) updateArmadaChartGroup(mgr armadaif.ArmadaManager, instance *av1.ArmadaChartGroup) (bool, error) {
	previousResource, updatedResource, err := mgr.UpdateResource(context.TODO())
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
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(instance.GetDependantResources()); err != nil {
			log.Error(err, "Failed to run update resource dependant resources")
			return false, err
		}
	}

	hrc := av1.HelmResourceCondition{
		Type:         av1.ConditionDeployed,
		Status:       av1.ConditionStatusTrue,
		Reason:       av1.ReasonUpdateSuccessful,
		Message:      "HardcodedMessage",
		ResourceName: updatedResource.GetName(),
	}
	instance.Status.SetCondition(hrc)
	instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// reconcileArmadaChartGroup reconciles the release with the cluster
func (r ChartGroupReconciler) reconcileArmadaChartGroup(mgr armadaif.ArmadaManager, instance *av1.ArmadaChartGroup) error {
	expectedResource, err := mgr.ReconcileResource(context.TODO())
	if err != nil {
		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionIrreconcilable,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonReconcileError,
			Message:      err.Error(),
			ResourceName: expectedResource.GetName(),
		}
		instance.Status.SetCondition(hrc)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)

	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(instance.GetDependantResources()); err != nil {
			log.Error(err, "Failed to run update resource dependant resources")
			return err
		}
	}
	return nil
}
