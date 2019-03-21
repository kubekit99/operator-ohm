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
	"reflect"
	"time"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	armadamgr "github.com/kubekit99/operator-ohm/armada-operator/pkg/armada"
	armadaif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	crthandler "sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	crtpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var acglog = logf.Log.WithName("acg-controller")

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
	// EnqueueRequestForObject enqueues a Request containing the Name and Namespace of the object
	// that is the source of the Event. (e.g. the created / deleted / updated objects Name and Namespace).
	err = c.Watch(&source.Kind{Type: &av1.ArmadaChartGroup{}}, &crthandler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	dependentPredicate := crtpredicate.Funcs{
		// We don't need to reconcile dependent resource creation events
		// because dependent resources are only ever created during
		// reconciliation. Another reconcile would be redundant.
		CreateFunc: func(e event.CreateEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			acglog.Info("CreateEvent. Filtering", "ArmadaChart", o.GetName(), "namespace", o.GetNamespace())
			return false
		},

		// Reconcile when a dependent resource is deleted so that it can be
		// recreated.
		DeleteFunc: func(e event.DeleteEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			acglog.Info("DeleteEvent. Triggering", "ArmadaChart", o.GetName(), "namespace", o.GetNamespace())
			return true
		},

		// Reconcile when a dependent resource is updated, so that it can
		// be patched back to the resource managed by the Helm release, if
		// necessary. Ignore updates that only change the status and
		// resourceVersion.
		UpdateFunc: func(e event.UpdateEvent) bool {
			old := e.ObjectOld.(*unstructured.Unstructured).DeepCopy()
			new := e.ObjectNew.(*unstructured.Unstructured).DeepCopy()

			delete(old.Object, "status")
			delete(new.Object, "status")
			old.SetResourceVersion("")
			new.SetResourceVersion("")

			if reflect.DeepEqual(old.Object, new.Object) {
				acglog.Info("UpdateEvent. Filtering", "ArmadaChart", new.GetName(), "namespace", new.GetNamespace())
				return false
			} else {
				acglog.Info("UpdateEvent. Triggering", "ArmadaChart", new.GetName(), "namespace", new.GetNamespace())
				return true
			}
		},
	}

	// Watch for changes to ArmadaChart and requeue the owner ArmadaChartGroup
	// EnqueueRequestForOwner enqueues Requests for the Owners of an object. E.g. the object
	// that created the object that was the source of the Event
	// IsController if set will only look at the first OwnerReference with Controller: true.
	act := av1.NewArmadaChartVersionKind("", "")
	err = c.Watch(&source.Kind{Type: act},
		&crthandler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &av1.ArmadaChartGroup{},
		},
		dependentPredicate)
	if err != nil {
		return err
	}

	// JEB: Will see later if we need to put the ownership between the backup/restore and the ChartGroup
	// err = c.Watch(&source.Kind{Type: &av1.ArmadaBackup{}}, &crthandler.EnqueueRequestForOwner{OwnerType: owner},
	// 	dependentPredicate)
	// err = c.Watch(&source.Kind{Type: &av1.ArmadaRestore{}}, &crthandler.EnqueueRequestForOwner{OwnerType: owner},
	// 	dependentPredicate)

	return nil
}

var _ reconcile.Reconciler = &ChartGroupReconciler{}

// ChartGroupReconciler reconciles a ArmadaChartGroup object
type ChartGroupReconciler struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client          client.Client
	scheme          *runtime.Scheme
	recorder        record.EventRecorder
	managerFactory  armadaif.ArmadaManagerFactory
	reconcilePeriod time.Duration
}

const (
	finalizerArmadaChartGroup = "uninstall-acg"
)

// Reconcile reads that state of the cluster for a ArmadaChartGroup object and
// makes changes based on the state read and what is in the ArmadaChartGroup.Spec
//
// Note: The Controller will requeue the Request to be processed again if the
// returned error is non-nil or Result.Requeue is true, otherwise upon
// completion it will remove the work from the queue.
func (r *ChartGroupReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reclog := acglog.WithValues("namespace", request.Namespace, "acg", request.Name)
	instance := &av1.ArmadaChartGroup{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if apierrors.IsNotFound(err) {
		// We are working asynchronously. By the time we receive the event,
		// the object is already gone
		return reconcile.Result{}, nil
	}
	if err != nil {
		reclog.Error(err, "Failed to lookup ArmadaChartGroup")
		return reconcile.Result{}, err
	}

	// AdminState POC begin
	// We will have to enhance the placement of this test to account
	// for kubectl apply where more than just the AdminState is changed
	if disabled := r.isReconcileDisabled(instance); disabled {
		return reconcile.Result{}, nil
	}
	// AdminState POC end

	mgr := r.managerFactory.NewArmadaChartGroupManager(instance)

	var shouldRequeue bool
	if shouldRequeue, err = r.updateFinalizers(instance); shouldRequeue {
		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	hrc := av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)

	if err := r.ensureSynced(mgr, instance); err != nil {
		if (!instance.IsDeleted()) {
			// TODO(jeb): Changed the behavior to stop only if we are not
			// in a delete phase.
			return reconcile.Result{}, err
		} 
	}

	switch {
	case instance.IsDeleted():
		if shouldRequeue, err = r.deleteArmadaChartGroup(mgr, instance); shouldRequeue {
			// Need to requeue because finalizer update does not change metadata.generation
			return reconcile.Result{Requeue: true}, err
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

	err = r.updateResourceStatus(instance)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// logAndRecordFailure adds a failure event to the recorder
func (r ChartGroupReconciler) logAndRecordFailure(instance *av1.ArmadaChartGroup, hrc *av1.HelmResourceCondition, err error) {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Error(err, fmt.Sprintf("%s. ErrorCondition", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

// logAndRecordSuccess adds a success event to the recorder
func (r ChartGroupReconciler) logAndRecordSuccess(instance *av1.ArmadaChartGroup, hrc *av1.HelmResourceCondition) {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Info(fmt.Sprintf("%s. SuccessCondition", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// updateResource updates the Resource object in the cluster
func (r ChartGroupReconciler) updateResource(o *av1.ArmadaChartGroup) error {
	return r.client.Update(context.TODO(), o)
}

// updateResourceStatus updates the the Status field of the Resource object in the cluster
func (r ChartGroupReconciler) updateResourceStatus(instance *av1.ArmadaChartGroup) error {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)

	helper := av1.HelmResourceConditionListHelper{Items: instance.Status.Conditions}
	instance.Status.Conditions = helper.InitIfEmpty()

	// JEB: Be sure to have update status subresources in the CRD.yaml
	// JEB: Look for kubebuilder subresources in the _types.go
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reclog.Error(err, "Failure to update ChartGroupStatus")
		// err = nil
	}

	return err
}

// ensureSynced checks that the ArmadaChartGroupManager is in sync with the cluster
func (r ChartGroupReconciler) ensureSynced(mgr armadaif.ArmadaChartGroupManager, instance *av1.ArmadaChartGroup) error {
	if err := mgr.Sync(context.TODO()); err != nil {
		hrc := av1.HelmResourceCondition{
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

// watchArmadaCharts updates all resources which are dependent on this one
func (r ChartGroupReconciler) watchArmadaCharts(instance *av1.ArmadaChartGroup, toWatchList *av1.ArmadaCharts) error {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Info("Adding Watch")

	errs := make([]error, 0)
	for _, toWatch := range (*toWatchList).List.Items {
		found := toWatch.FromArmadaChart()
		err := r.client.Get(context.TODO(), types.NamespacedName{Name: found.GetName(), Namespace: found.GetNamespace()}, found)
		if err == nil {
			if err1 := controllerutil.SetControllerReference(instance, found, r.scheme); err1 != nil {
				reclog.Error(err1, "Can't get ownership of ArmadaChart", "name", found.GetName())
				errs = append(errs, err1)
				continue
			}
			if err2 := r.client.Update(context.TODO(), found); err2 != nil {
				reclog.Error(err2, "Can't get ownership of ArmadaChart", "name", found.GetName())
				errs = append(errs, err2)
				continue
			}
			reclog.Info("Added ownership of ArmadaChart", "name", found.GetName())
		} else {
			reclog.Error(err, "Can't get ownership of ArmadaChart", "name", found.GetName())
			errs = append(errs, err)
		}
	}

	return nil
}

// deleteArmadaChartGroup deletes an instance of an ArmadaChartGroup. It returns true if the reconciler should be re-enqueueed
func (r ChartGroupReconciler) deleteArmadaChartGroup(mgr armadaif.ArmadaChartGroupManager, instance *av1.ArmadaChartGroup) (bool, error) {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Info("Deleting")

	pendingFinalizers := instance.GetFinalizers()
	if !contains(pendingFinalizers, finalizerArmadaChartGroup) {
		reclog.Info("ChartGroup is terminated, skipping reconciliation")
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
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err == armadaif.ErrNotFound {
		reclog.Info("Charts are already deleted, removing finalizer")
	} else {
		hrc := av1.HelmResourceCondition{
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
		if pendingFinalizer != finalizerArmadaChartGroup {
			finalizers = append(finalizers, pendingFinalizer)
		}
	}
	instance.SetFinalizers(finalizers)
	err = r.updateResource(instance)

	// Need to requeue because finalizer update does not change metadata.generation
	return true, err
}

// updateArmadaChartGroup attempts to update instance. It returns true if the reconciler should be re-enqueueed
func (r ChartGroupReconciler) updateArmadaChartGroup(mgr armadaif.ArmadaChartGroupManager, instance *av1.ArmadaChartGroup) (bool, error) {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Info("Updating")

	_, updatedResource, err := mgr.UpdateResource(context.TODO())

	// TODO(jeb): Behavior is flacky here. err != nil means updatedResource is nil
	// Watch for panic exception if UpdateResource behavior is modified
	if err != nil {
		hrc := av1.HelmResourceCondition{
			Type:         av1.ConditionFailed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUpdateError,
			Message:      err.Error(),
			ResourceName: "",
		}
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return false, err
	}
	instance.Status.RemoveCondition(av1.ConditionFailed)

	if err := r.watchArmadaCharts(instance, updatedResource); err != nil {
		return false, err
	}

	hrc := av1.HelmResourceCondition{
		Type:         av1.ConditionDeployed,
		Status:       av1.ConditionStatusTrue,
		Reason:       av1.ReasonUpdateSuccessful,
		Message:      "HardcodedMessage",
		ResourceName: updatedResource.GetName(),
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// reconcileArmadaChartGroup reconciles the release with the cluster
func (r ChartGroupReconciler) reconcileArmadaChartGroup(mgr armadaif.ArmadaChartGroupManager, instance *av1.ArmadaChartGroup) error {
	reclog := acglog.WithValues("namespace", instance.Namespace, "acg", instance.Name)
	reclog.Info("Reconciling ArmadaChartGroup and ArmadaChartList")

	expectedResource, err := mgr.ReconcileResource(context.TODO())
	if err != nil {
		hrc := av1.HelmResourceCondition{
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
	err = r.watchArmadaCharts(instance, expectedResource)
	return err
}

// isReconcileDisabled
func (r ChartGroupReconciler) isReconcileDisabled(instance *av1.ArmadaChartGroup) bool {
	// JEB: Not sure if we need to add this new ConditionEnabled
	// or we can just used the ConditionInitialized
	if instance.IsDisabled() {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusFalse,
			Reason: "ChartGroup is disabled",
		}
		r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		_ = r.updateResourceStatus(instance)
		return true
	} else {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusTrue,
			Reason: "ChartGroup is enabled",
		}
		r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		return false
	}
}
