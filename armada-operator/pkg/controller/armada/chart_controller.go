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

// Add creates a new ArmadaChart Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddArmadaChartController(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	r := &ArmadaChartReconciler{
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
	if racr, isArmadaChartReconciler := r.(*ArmadaChartReconciler); isArmadaChartReconciler {
		owner := av1.NewArmadaChartVersionKind("", "")
		racr.depResourceWatchUpdater = services.BuildDependantResourceWatchUpdater(mgr, owner, c)
	} else if rrf, isReconcileFunc := r.(*reconcile.Func); isReconcileFunc {
		log.Info("UnitTests", "ReconfileFunc", rrf)
		// JEB: This the wrapper used during the unit tests
		// owner := av1.NewArmadaChartVersionKind("", "")
		// rrf.inner.depResourceWatchUpdater = services.BuildDependantResourceWatchUpdater(mgr, owner, c)
	}

	return nil
}

var _ reconcile.Reconciler = &ArmadaChartReconciler{}

// ArmadaChartReconciler reconciles custom resources as Helm releases.
type ArmadaChartReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          services.HelmManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater services.DependantResourceWatchUpdater
}

const (
	finalizerArmadaChart = "uninstall-helm-release"
)

// Reconcile reads that state of the cluster for a ArmadaChart object and makes changes based on the state read
// and what is in the ArmadaChart.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ArmadaChartReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	log := log.WithValues(
		"namespace", request.Namespace,
		"name", request.Name,
	)

	instance, err := r.getArmadaChartInstance(request)
	if err != nil {
		log.Error(err, "Failed to lookup resource")
		return reconcile.Result{}, err
	}

	manager := r.managerFactory.NewArmadaChartTillerManager(instance)
	log = log.WithValues("resource", manager.ReleaseName())

	pendingFinalizers := instance.GetFinalizers()
	if !instance.IsDeleted() && !contains(pendingFinalizers, finalizerArmadaChart) {
		finalizers := append(pendingFinalizers, finalizerArmadaChart)
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	hrc := getConditionInitialized()
	instance.Status.SetCondition(hrc)
	instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)

	if err := r.ensureSynced(manager, instance); err != nil {
		return reconcile.Result{}, err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)

	if instance.IsDeleted() {
		if !contains(pendingFinalizers, finalizerArmadaChart) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledResource, err := manager.UninstallRelease(context.TODO())
		if err != nil && err != services.ErrNotFound {
			hrc := getConditionUninstallError(uninstalledResource, err.Error())
			instance.Status.SetCondition(hrc)
			instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance)
			return reconcile.Result{}, err
		}
		instance.Status.RemoveCondition(av1.ConditionFailed)

		if err == services.ErrNotFound {
			log.Info("Resource not found, removing finalizer")
		} else {
			hrc := getConditionUninstallSuccessful()
			instance.Status.SetCondition(hrc)
			instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
			r.logAndRecordSuccess(instance, &hrc)
		}
		if err := r.updateResourceStatus(instance); err != nil {
			return reconcile.Result{}, err
		}

		finalizers := []string{}
		for _, pendingFinalizer := range pendingFinalizers {
			if pendingFinalizer != finalizerArmadaChart {
				finalizers = append(finalizers, pendingFinalizer)
			}
		}
		instance.SetFinalizers(finalizers)
		err = r.updateResource(instance)

		// Need to requeue because finalizer update does not change metadata.generation
		return reconcile.Result{Requeue: true}, err
	}

	if !manager.IsInstalled() {
		installedResource, err := manager.InstallRelease(context.TODO())
		if err != nil {
			hrc := getConditionInstallError(err.Error())
			instance.Status.SetCondition(hrc)
			instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance)
			return reconcile.Result{}, err
		}
		instance.Status.RemoveCondition(av1.ConditionFailed)

		if r.depResourceWatchUpdater != nil {
			if err := r.depResourceWatchUpdater(helmmgr.GetDependantResources(installedResource)); err != nil {
				log.Error(err, "Failed to run update resource dependant resources")
				return reconcile.Result{}, err
			}
		}

		hrc := getConditionInstallSuccess(installedResource)
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)

		err = r.updateResourceStatus(instance)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	if manager.IsUpdateRequired() {
		previousResource, updatedResource, err := manager.UpdateRelease(context.TODO())
		if previousResource != nil && updatedResource != nil {
			log.Info("UpdateRelease", "Previous", previousResource.GetName(), "Updated", updatedResource.GetName())
		}
		if err != nil {
			hrc := getConditionUpdateError(updatedResource, err.Error())
			instance.Status.SetCondition(hrc)
			instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
			r.logAndRecordFailure(instance, &hrc, err)

			_ = r.updateResourceStatus(instance)
			return reconcile.Result{}, err
		}
		instance.Status.RemoveCondition(av1.ConditionFailed)

		if r.depResourceWatchUpdater != nil {
			if err := r.depResourceWatchUpdater(helmmgr.GetDependantResources(updatedResource)); err != nil {
				log.Error(err, "Failed to run update resource dependant resources")
				return reconcile.Result{}, err
			}
		}

		hrc := getConditionUpdateSuccessful(updatedResource)
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordSuccess(instance, &hrc)

		err = r.updateResourceStatus(instance)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	expectedResource, err := manager.ReconcileRelease(context.TODO())
	if err != nil {
		hrc := getConditionIrreconcilable(expectedResource, err.Error())
		instance.Status.SetCondition(hrc)
		r.logAndRecordFailure(instance, &hrc, err)

		_ = r.updateResourceStatus(instance)
		return reconcile.Result{}, err
	}
	instance.Status.RemoveCondition(av1.ConditionIrreconcilable)

	if r.depResourceWatchUpdater != nil {
		if err := r.depResourceWatchUpdater(helmmgr.GetDependantResources(expectedResource)); err != nil {
			log.Error(err, "Failed to run update resource dependant resources")
			return reconcile.Result{}, err
		}
	}

	log.Info("Reconciled resource")
	err = r.updateResourceStatus(instance)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
}

func (r *ArmadaChartReconciler) getArmadaChartInstance(request reconcile.Request) (*av1.ArmadaChart, error) {
	instance := &av1.ArmadaChart{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)

	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if apierrors.IsNotFound(err) {
		return &av1.ArmadaChart{}, err
	}
	return instance, nil
}

// Add a success event to the recorder
func (r ArmadaChartReconciler) logAndRecordFailure(instance *av1.ArmadaChart, hrc *av1.HelmResourceCondition, err error) {
	log.Error(err, fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeWarning, hrc.Type.String(), hrc.Reason.String())
}

func (r ArmadaChartReconciler) logAndRecordSuccess(instance *av1.ArmadaChart, hrc *av1.HelmResourceCondition) {
	log.Info(fmt.Sprintf("%s", hrc.Type.String()))
	r.recorder.Event(instance, corev1.EventTypeNormal, hrc.Type.String(), hrc.Reason.String())
}

// Update the Resource object
func (r ArmadaChartReconciler) updateResource(o *av1.ArmadaChart) error {
	return r.client.Update(context.TODO(), o)
}

// Update the Status field in the CRD
func (r ArmadaChartReconciler) updateResourceStatus(instance *av1.ArmadaChart) error {
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

func (r ArmadaChartReconciler) ensureSynced(mgr services.HelmManager, instance *av1.ArmadaChart) error {
	if err := mgr.Sync(context.TODO()); err != nil {
		hrc := getConditionIrreconcilable(nil, err.Error())
		instance.Status.SetCondition(hrc)
		instance.Status.ComputeActualState(&hrc, instance.Spec.TargetState)
		r.logAndRecordFailure(instance, &hrc, err)
		_ = r.updateResourceStatus(instance)
		return err
	}
	return nil
}

func getConditionInitialized() av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	}
}

func getConditionIrreconcilable(resource *release.Release, message string) av1.HelmResourceCondition {
	hrc := av1.HelmResourceCondition{
		Type:    av1.ConditionIrreconcilable,
		Status:  av1.ConditionStatusTrue,
		Reason:  av1.ReasonReconcileError,
		Message: message,
	}

	if resource != nil {
		hrc.ResourceName = resource.GetName()
	}

	return hrc
}

func getConditionUninstallError(resource *release.Release, message string) av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:         av1.ConditionFailed,
		Status:       av1.ConditionStatusTrue,
		Reason:       av1.ReasonUninstallError,
		Message:      message,
		ResourceName: resource.GetName(),
	}
}

func getConditionUninstallSuccessful() av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:   av1.ConditionDeployed,
		Status: av1.ConditionStatusFalse,
		Reason: av1.ReasonUninstallSuccessful,
	}
}

func getConditionInstallError(message string) av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:    av1.ConditionFailed,
		Status:  av1.ConditionStatusTrue,
		Reason:  av1.ReasonInstallError,
		Message: message,
	}
}

func getConditionInstallSuccess(resource *release.Release) av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonInstallSuccessful,
		Message:         resource.GetInfo().GetStatus().GetNotes(),
		ResourceName:    resource.GetName(),
		ResourceVersion: resource.GetVersion(),
	}
}

func getConditionUpdateError(resource *release.Release, message string) av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:         av1.ConditionFailed,
		Status:       av1.ConditionStatusTrue,
		Reason:       av1.ReasonUpdateError,
		Message:      message,
		ResourceName: resource.GetName(),
	}
}

func getConditionUpdateSuccessful(resource *release.Release) av1.HelmResourceCondition {
	return av1.HelmResourceCondition{
		Type:            av1.ConditionDeployed,
		Status:          av1.ConditionStatusTrue,
		Reason:          av1.ReasonUpdateSuccessful,
		Message:         resource.GetInfo().GetStatus().GetNotes(),
		ResourceName:    resource.GetName(),
		ResourceVersion: resource.GetVersion(),
	}
}
