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

package armada

import (
	"context"
	"fmt"
	"time"

	helmmgr "github.com/kubekit99/operator-ohm/armada-operator/pkg/helm"
	services "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"k8s.io/helm/pkg/proto/hapi/release"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// AddArmadaChartController creates a new ArmadaChart Controller and adds it to
// the Manager. The Manager will set fields on the Controller and Start it when
// the Manager is Started.
func AddArmadaChartController(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	r := &ChartReconciler{
		client:         mgr.GetClient(),
		scheme:         mgr.GetScheme(),
		recorder:       mgr.GetRecorder("act-recorder"),
		managerFactory: helmmgr.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
	}
	return r
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {

	// Create a new controller
	c, err := controller.New("act-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArmadaChart
	err = c.Watch(&source.Kind{Type: &av1.ArmadaChart{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner ArmadaChart
	if racr, isChartReconciler := r.(*ChartReconciler); isChartReconciler {
		owner := av1.NewArmadaChartVersionKind("", "")
		racr.depResourceWatchUpdater = services.BuildDependentResourceWatchUpdater(mgr, owner, c)
	} else if rrf, isReconcileFunc := r.(*reconcile.Func); isReconcileFunc {
		log.Info("UnitTests", "ReconfileFunc", rrf)
		// JEB: This the wrapper used during the unit tests
		// owner := av1.NewArmadaChartVersionKind("", "")
		// rrf.inner.depResourceWatchUpdater = services.BuildDependentResourceWatchUpdater(mgr, owner, c)
	}

	return nil
}

var _ reconcile.Reconciler = &ChartReconciler{}

// ChartReconciler reconciles custom resources as Helm releases.
type ChartReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          services.HelmManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater services.DependentResourceWatchUpdater
}

const (
	finalizerArmadaChart = "uninstall-helm-release"
)

// Reconcile reads that state of the cluster for an ArmadaChart object and
// makes changes based on the state read and what is in the ArmadaChart.Spec
//
// Note: The Controller will requeue the Request to be processed again if the
// returned error is non-nil or Result.Requeue is true, otherwise upon
// completion it will remove the work from the queue.
func (r *ChartReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	log := log.WithValues(
		"namespace", request.Namespace,
		"name", request.Name,
	)

	instance, err := r.getArmadaChartInstance(request)
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

	mgr := r.managerFactory.NewArmadaChartTillerManager(instance)
	log = log.WithValues("resource", mgr.ReleaseName())

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
		return reconcile.Result{}, err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)

	switch {
	case instance.IsDeleted():
		if shouldRequeue, err = r.deleteArmadaChart(mgr, instance); shouldRequeue {
			// Need to requeue because finalizer update does not change metadata.generation
			return reconcile.Result{Requeue: true}, err
		}
		return reconcile.Result{}, err
	case !mgr.IsInstalled():
		if shouldRequeue, err = r.installArmadaChart(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	case mgr.IsUpdateRequired():
		if shouldRequeue, err = r.updateArmadaChart(mgr, instance); shouldRequeue {
			return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
		}
		return reconcile.Result{}, err
	}

	if err := r.reconcileArmadaChart(mgr, instance); err != nil {
		return reconcile.Result{}, err
	}

	log.Info("Reconciled resource")
	err = r.updateResourceStatus(instance)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

// getArmadaChartInstance gets an instance of the ArmadaChart which has the
// name and namespace specified in the request
func (r *ChartReconciler) getArmadaChartInstance(request reconcile.Request) (*av1.ArmadaChart, error) {
	instance := &av1.ArmadaChart{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if apierrors.IsNotFound(err) {
		return &av1.ArmadaChart{}, err
	}
	return instance, nil
}

// logAndRecordFailure adds a failure event to the recorder
func (r ChartReconciler) logAndRecordFailure(instance *av1.ArmadaChart, hrc *av1.HelmResourceCondition, err error) {
	log.Error(err, fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

// logAndRecordSuccess adds a success event to the recorder
func (r ChartReconciler) logAndRecordSuccess(instance *av1.ArmadaChart, hrc *av1.HelmResourceCondition) {
	log.Info(fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// updateResource updates the Resource object in the cluster
func (r ChartReconciler) updateResource(instance *av1.ArmadaChart) error {
	return r.client.Update(context.TODO(), instance)
}

// updateResourceStatus updates the the Status field of the Resource object in the cluster
func (r ChartReconciler) updateResourceStatus(instance *av1.ArmadaChart) error {
	reqLogger := log.WithValues("ArmadaChart.Namespace", instance.Namespace, "ArmadaChart.Name", instance.Name)

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

// ensureSynced checks that the HelmManager is in sync with the cluster
func (r ChartReconciler) ensureSynced(mgr services.HelmManager, instance *av1.ArmadaChart) error {
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
	return nil
}

// updateFinalizers asserts that the finalizers match what is expected based on
// whether the instance is currently being deleted or not. It returns true if
// the finalizers were changed, false otherwise
func (r ChartReconciler) updateFinalizers(instance *av1.ArmadaChart) (bool, error) {
	pendingFinalizers := instance.GetFinalizers()
	if !instance.IsDeleted() && !contains(pendingFinalizers, finalizerArmadaChart) {
		finalizers := append(pendingFinalizers, finalizerArmadaChart)
		instance.SetFinalizers(finalizers)
		err := r.updateResource(instance)

		return true, err
	}
	return false, nil
}

// updateDependentResources updates all resources which are dependent on this one
func (r ChartReconciler) updateDependentResources(resource *release.Release) error {
	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(helmmgr.GetDependentResources(resource)); err != nil {
			log.Error(err, "Failed to run update resource dependent resources")
			return err
		}
	}
	return nil
}

// deleteArmadaChart deletes an instance of an ArmadaChart. It returns true if the reconciler should be re-enqueueed
func (r ChartReconciler) deleteArmadaChart(mgr services.HelmManager, instance *av1.ArmadaChart) (bool, error) {
	pendingFinalizers := instance.GetFinalizers()
	if !contains(pendingFinalizers, finalizerArmadaChart) {
		log.Info("Resource is terminated, skipping reconciliation")
		return false, nil
	}

	uninstalledResource, err := mgr.UninstallRelease(context.TODO())
	if err != nil && err != services.ErrNotFound {
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

	if err == services.ErrNotFound {
		log.Info("Resource not found, removing finalizer")
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
		if pendingFinalizer != finalizerArmadaChart {
			finalizers = append(finalizers, pendingFinalizer)
		}
	}
	instance.SetFinalizers(finalizers)
	err = r.updateResource(instance)

	return true, err
}

// installArmadaChart attempts to install instance. It returns true if the reconciler should be re-enqueueed
func (r ChartReconciler) installArmadaChart(mgr services.HelmManager, instance *av1.ArmadaChart) (bool, error) {
	installedResource, err := mgr.InstallRelease(context.TODO())
	if err != nil {
		hrc := av1.HelmResourceCondition{
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

	if err := r.updateDependentResources(installedResource); err != nil {
		return false, err
	}

	hrc := av1.HelmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonInstallSuccessful,
		Message:         installedResource.GetInfo().GetStatus().GetNotes(),
		ResourceName:    installedResource.GetName(),
		ResourceVersion: installedResource.GetVersion(),
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// updateArmadaChart attempts to update instance. It returns true if the reconciler should be re-enqueueed
func (r ChartReconciler) updateArmadaChart(mgr services.HelmManager, instance *av1.ArmadaChart) (bool, error) {
	previousResource, updatedResource, err := mgr.UpdateRelease(context.TODO())
	if previousResource != nil && updatedResource != nil {
		log.Info("UpdateRelease", "Previous", previousResource.GetName(), "Updated", updatedResource.GetName())
	}
	if err != nil {
		hrc := av1.HelmResourceCondition{
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

	if err := r.updateDependentResources(updatedResource); err != nil {
		return false, err
	}

	hrc := av1.HelmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonUpdateSuccessful,
		Message:         updatedResource.GetInfo().GetStatus().GetNotes(),
		ResourceName:    updatedResource.GetName(),
		ResourceVersion: updatedResource.GetVersion(),
	}
	instance.Status.SetCondition(hrc, instance.Spec.TargetState)
	r.logAndRecordSuccess(instance, &hrc)

	err = r.updateResourceStatus(instance)
	return true, err
}

// reconcileArmadaChart reconciles the release with the cluster
func (r ChartReconciler) reconcileArmadaChart(mgr services.HelmManager, instance *av1.ArmadaChart) error {
	expectedResource, err := mgr.ReconcileRelease(context.TODO())
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
	err = r.updateDependentResources(expectedResource)
	return err
}

// isReconcileDisabled
func (r ChartReconciler) isReconcileDisabled(instance *av1.ArmadaChart) bool {
	// JEB: Not sure if we need to add this new ConditionEnabled
	// or we can just used the ConditionInitialized
	if instance.IsDisabled() {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusFalse,
			Reason: "Chart is disabled",
		}
		r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		_ = r.updateResourceStatus(instance)
		return true
	} else {
		hrc := av1.HelmResourceCondition{
			Type:   av1.ConditionEnabled,
			Status: av1.ConditionStatusTrue,
			Reason: "Chart is enabled",
		}
		r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
		instance.Status.SetCondition(hrc, instance.Spec.TargetState)
		return false
	}
}
