// Copyright 2019 The Openstack-Service-Lifecyle Authors
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

package osphases

import (
	"context"
	"fmt"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
	trafficdrainphasemgr "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/osphases"
	services "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"sigs.k8s.io/controller-runtime/pkg/controller"
	crthandler "sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var trafficdrainphaselog = logf.Log.WithName("trafficdrainphase-controller")

// AddTrafficDrainPhaseController creates a new TrafficDrainPhase Controller and adds it to
// the Manager. The Manager will set fields on the Controller and Start it when
// the Manager is Started.
func AddTrafficDrainPhaseController(mgr manager.Manager) error {
	return addTrafficDrainPhase(mgr, newTrafficDrainPhaseReconciler(mgr))
}

// newTrafficDrainPhaseReconciler returns a new reconcile.Reconciler
func newTrafficDrainPhaseReconciler(mgr manager.Manager) reconcile.Reconciler {
	r := &TrafficDrainPhaseReconciler{
		PhaseReconciler: PhaseReconciler{
			client:         mgr.GetClient(),
			scheme:         mgr.GetScheme(),
			recorder:       mgr.GetRecorder("trafficdrainphase-recorder"),
			managerFactory: trafficdrainphasemgr.NewManagerFactory(mgr),
			// reconcilePeriod: flags.ReconcilePeriod,
		},
	}
	return r
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func addTrafficDrainPhase(mgr manager.Manager, r reconcile.Reconciler) error {

	// Create a new controller
	c, err := controller.New("trafficdrainphase-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource TrafficDrainPhase
	// EnqueueRequestForObject enqueues a Request containing the Name and Namespace of the object
	// that is the source of the Event. (e.g. the created / deleted / updated objects Name and Namespace).
	err = c.Watch(&source.Kind{Type: &av1.TrafficDrainPhase{}}, &crthandler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource (described in the yaml files) and requeue the owner TrafficDrainPhase
	// EnqueueRequestForOwner enqueues Requests for the Owners of an object. E.g. the object
	// that created the object that was the source of the Event
	if racr, isTrafficDrainPhaseReconciler := r.(*TrafficDrainPhaseReconciler); isTrafficDrainPhaseReconciler {
		// The enqueueRequestForOwner is not actually done here since we don't know yet the
		// content of the yaml files. The tools wait for the yaml files to be parse. The manager
		// then add the "OwnerReference" to the content of the yaml files. It then invokes the EnqueueRequestForOwner
		owner := av1.NewTrafficDrainPhaseVersionKind("", "")
		dependentPredicate := racr.BuildDependentPredicate()
		racr.depResourceWatchUpdater = services.BuildDependentResourceWatchUpdater(mgr, owner, c, *dependentPredicate)
	} else if rrf, isReconcileFunc := r.(*reconcile.Func); isReconcileFunc {
		// Unit test issue
		log.Info("UnitTests", "ReconfileFunc", rrf)
	}

	return nil
}

var _ reconcile.Reconciler = &TrafficDrainPhaseReconciler{}

// TrafficDrainPhaseReconciler reconciles custom resources as Argo workflows.
type TrafficDrainPhaseReconciler struct {
	PhaseReconciler
}

const (
	finalizerTrafficDrainPhase = "uninstall-trafficdrainphase-resource"
)

// Reconcile reads that state of the cluster for an TrafficDrainPhase object and
// makes changes based on the state read and what is in the TrafficDrainPhase.Spec
//
// Note: The Controller will requeue the Request to be processed again if the
// returned error is non-nil or Result.Requeue is true, otherwise upon
// completion it will remove the work from the queue.
func (r *TrafficDrainPhaseReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reclog := trafficdrainphaselog.WithValues("namespace", request.Namespace, "trafficdrainphase", request.Name)
	reclog.Info("Received a request")

	instance := &av1.TrafficDrainPhase{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	instance.Init()

	if apierrors.IsNotFound(err) {
		// We are working asynchronously. By the time we receive the event,
		// the object could already be gone
		return reconcile.Result{}, nil
	}

	if err != nil {
		reclog.Error(err, "Failed to lookup TrafficDrainPhase")
		return reconcile.Result{}, err
	}

	mgr := r.managerFactory.NewTrafficDrainPhaseManager(instance)
	reclog = reclog.WithValues("trafficdrainphase", mgr.ResourceName())

	var shouldRequeue bool
	if shouldRequeue, err = r.updateFinalizers(instance); shouldRequeue {
		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	if err := r.ensureSynced(mgr, instance); err != nil {
		if !instance.IsDeleted() {
			// TODO(jeb): Changed the behavior to stop only if we are not
			// in a delete phase.
			return reconcile.Result{}, err
		}
	}

	if instance.IsDeleted() {
		if shouldRequeue, err = r.deleteTrafficDrainPhase(mgr, instance); shouldRequeue {
			// Need to requeue because finalizer update does not change metadata.generation
			return reconcile.Result{Requeue: true}, err
		}
		return reconcile.Result{}, err
	}

	if instance.IsTargetStateUninitialized() {
		reclog.Info("TargetState unitialized; skipping")
		err = r.updateResource(instance)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.client.Status().Update(context.TODO(), instance)
		return reconcile.Result{}, err
	}

	hrc := av1.LcmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)

	switch {
	case !mgr.IsInstalled():
		if shouldRequeue, err = r.installTrafficDrainPhase(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	case mgr.IsUpdateRequired():
		if shouldRequeue, err = r.updateTrafficDrainPhase(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	}

	if err := r.reconcileTrafficDrainPhase(mgr, instance); err != nil {
		return reconcile.Result{}, err
	}

	reclog.Info("Reconciled TrafficDrainPhase")
	err = r.updateResourceStatus(instance)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// logAndRecordFailure adds a failure event to the recorder
func (r TrafficDrainPhaseReconciler) logAndRecordFailure(instance *av1.TrafficDrainPhase, hrc *av1.LcmResourceCondition, err error) {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Error(err, fmt.Sprintf("%s. ErrorCondition", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

// logAndRecordSuccess adds a success event to the recorder
func (r TrafficDrainPhaseReconciler) logAndRecordSuccess(instance *av1.TrafficDrainPhase, hrc *av1.LcmResourceCondition) {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Info(fmt.Sprintf("%s. SuccessCondition", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// updateResource updates the Resource object in the cluster
func (r TrafficDrainPhaseReconciler) updateResource(instance *av1.TrafficDrainPhase) error {
	return r.client.Update(context.TODO(), instance)
}

// updateResourceStatus updates the the Status field of the Resource object in the cluster
func (r TrafficDrainPhaseReconciler) updateResourceStatus(instance *av1.TrafficDrainPhase) error {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)

	helper := av1.LcmResourceConditionListHelper{Items: instance.Status.Conditions}
	instance.Status.Conditions = helper.InitIfEmpty()

	// JEB: Be sure to have update status subresources in the CRD.yaml
	// JEB: Look for kubebuilder subresources in the _types.go
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reclog.Error(err, "Failure to update status. Ignoring")
		err = nil
	}

	return err
}

// ensureSynced checks that the TrafficDrainPhaseManager is in sync with the cluster
func (r TrafficDrainPhaseReconciler) ensureSynced(mgr services.TrafficDrainPhaseManager, instance *av1.TrafficDrainPhase) error {
	if err := mgr.Sync(context.TODO()); err != nil {
		hrc := av1.LcmResourceCondition{
			Type:    av1.ConditionIrreconcilable,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonReconcileError,
			Message: err.Error(),
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)
		_ = r.updateResourceStatus(instance)
		return err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)
	return nil
}

// updateFinalizers asserts that the finalizers match what is expected based on
// whether the instance is currently being deleted or not. It returns true if
// the finalizers were changed, false otherwise
func (r TrafficDrainPhaseReconciler) updateFinalizers(instance *av1.TrafficDrainPhase) (bool, error) {
	pendingFinalizers := instance.GetFinalizers()
	if !instance.IsDeleted() && !r.contains(pendingFinalizers, finalizerTrafficDrainPhase) {
		finalizers := append(pendingFinalizers, finalizerTrafficDrainPhase)
		instance.SetFinalizers(finalizers)
		err := r.updateResource(instance)

		return true, err
	}
	return false, nil
}

// watchDependentResources updates all resources which are dependent on this one
func (r TrafficDrainPhaseReconciler) watchDependentResources(resource *av1.SubResourceList) error {
	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(resource.GetDependentResources()); err != nil {
			return err
		}
	}
	return nil
}

// deleteTrafficDrainPhase deletes an instance of an TrafficDrainPhase. It returns true if the reconciler should be re-enqueueed
func (r TrafficDrainPhaseReconciler) deleteTrafficDrainPhase(mgr services.TrafficDrainPhaseManager, instance *av1.TrafficDrainPhase) (bool, error) {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Info("Deleting")

	pendingFinalizers := instance.GetFinalizers()
	if !r.contains(pendingFinalizers, finalizerTrafficDrainPhase) {
		reclog.Info("TrafficDrainPhase is terminated, skipping reconciliation")
		return false, nil
	}

	uninstalledResource, err := mgr.UninstallResource(context.TODO())
	if err != nil && err != services.ErrNotFound {
		hrc := av1.LcmResourceCondition{
			Type:         av1.ConditionFailed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUninstallError,
			Message:      err.Error(),
			ResourceName: uninstalledResource.GetName(),
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err == services.ErrNotFound {
		reclog.Info("Resource already uninstalled, Removing finalizer")
	} else {
		hrc := av1.LcmResourceCondition{
			Type:   av1.ConditionDeployed,
			Status: av1.ConditionStatusFalse,
			Reason: av1.ReasonUninstallSuccessful,
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)
	}
	if err := r.updateResourceStatus(instance); err != nil {
		return false, err
	}

	finalizers := []string{}
	for _, pendingFinalizer := range pendingFinalizers {
		if pendingFinalizer != finalizerTrafficDrainPhase {
			finalizers = append(finalizers, pendingFinalizer)
		}
	}
	instance.SetFinalizers(finalizers)
	err = r.updateResource(instance)

	return true, err
}

// installTrafficDrainPhase attempts to install instance. It returns true if the reconciler should be re-enqueueed
func (r TrafficDrainPhaseReconciler) installTrafficDrainPhase(mgr services.TrafficDrainPhaseManager, instance *av1.TrafficDrainPhase) (bool, error) {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Info("Installing`")

	installedResource, err := mgr.InstallResource(context.TODO())
	if err != nil {
		hrc := av1.LcmResourceCondition{
			Type:    av1.ConditionFailed,
			Status:  av1.ConditionStatusTrue,
			Reason:  av1.ReasonInstallError,
			Message: err.Error(),
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err := r.watchDependentResources(installedResource); err != nil {
		reclog.Error(err, "Failed to update watch on dependent resources")
		return false, err
	}

	hrc := av1.LcmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonInstallSuccessful,
		Message:         installedResource.GetNotes(),
		ResourceName:    installedResource.GetName(),
		ResourceVersion: installedResource.GetVersion(),
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// updateTrafficDrainPhase attempts to update instance. It returns true if the reconciler should be re-enqueueed
func (r TrafficDrainPhaseReconciler) updateTrafficDrainPhase(mgr services.TrafficDrainPhaseManager, instance *av1.TrafficDrainPhase) (bool, error) {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Info("Updating`")

	previousResource, updatedResource, err := mgr.UpdateResource(context.TODO())
	if previousResource != nil && updatedResource != nil {
		reclog.Info("UpdateResource", "Previous", previousResource.GetName(), "Updated", updatedResource.GetName())
	}
	if err != nil {
		hrc := av1.LcmResourceCondition{
			Type:         av1.ConditionFailed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUpdateError,
			Message:      err.Error(),
			ResourceName: updatedResource.GetName(),
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err := r.watchDependentResources(updatedResource); err != nil {
		reclog.Error(err, "Failed to update watch on dependent resources")
		return false, err
	}

	hrc := av1.LcmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonUpdateSuccessful,
		Message:         updatedResource.GetNotes(),
		ResourceName:    updatedResource.GetName(),
		ResourceVersion: updatedResource.GetVersion(),
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// reconcileTrafficDrainPhase reconciles the yaml files with the cluster
func (r TrafficDrainPhaseReconciler) reconcileTrafficDrainPhase(mgr services.TrafficDrainPhaseManager, instance *av1.TrafficDrainPhase) error {
	reclog := trafficdrainphaselog.WithValues("namespace", instance.Namespace, "trafficdrainphase", instance.Name)
	reclog.Info("Reconciling TrafficDrainPhase and LcmResource")

	expectedResource, err := mgr.ReconcileResource(context.TODO())
	if err != nil {
		hrc := av1.LcmResourceCondition{
			Type:         av1.ConditionIrreconcilable,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonReconcileError,
			Message:      err.Error(),
			ResourceName: expectedResource.GetName(),
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)
	if err := r.watchDependentResources(expectedResource); err != nil {
		reclog.Error(err, "Failed to update watch on dependent resources")
		return err
	}
	return nil
}
