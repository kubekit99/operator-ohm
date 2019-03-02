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
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	// "k8s.io/apimachinery/pkg/types"

	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	// "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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
		recorder:       mgr.GetRecorder("act-recorder"),
		managerFactory: armadamgr.NewManagerFactory(mgr),
		// reconcilePeriod: flags.ReconcilePeriod,
		watchDependentResources: true,
	}

	// Create a new controller
	c, err := controller.New("act-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArmadaManifest
	err = c.Watch(&source.Kind{Type: &av1.ArmadaManifest{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ArmadaManifest
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &av1.ArmadaManifest{},
	})
	if err != nil {
		return err
	}

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
	watchDependentResources bool
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

	manager := r.managerFactory.NewArmadaManifestManager(instance)
	// spec := instance.Spec
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

	status.SetCondition(av1.HelmResourceCondition{
		Type:   av1.ConditionInitialized,
		Status: av1.ConditionStatusTrue,
	})

	if err := manager.Sync(context.TODO()); err != nil {
		log.Error(err, "Failed to sync resource")
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
		if !contains(pendingFinalizers, finalizerArmadaManifest) {
			log.Info("Resource is terminated, skipping reconciliation")
			return reconcile.Result{}, nil
		}

		uninstalledResource, err := manager.UninstallResource(context.TODO())
		if err != nil && err != armadaif.ErrNotFound {
			log.Error(err, "Failed to uninstall resource")
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

		if err == armadaif.ErrNotFound {
			log.Info("Resource not found, removing finalizer")
		} else {
			r.recorder.Event(instance, corev1.EventTypeWarning, "DeletionFailure", fmt.Sprintf("Uninstalled Resource %s", uninstalledResource.GetName()))
			log.Info("Uninstalled resource", "resourceName", uninstalledResource.GetName())
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
			log.Error(err, "Failed to install resource")
			r.recorder.Event(instance, corev1.EventTypeWarning, "InstallationFailure", fmt.Sprintf("Installed Resource %s", installedResource.GetName()))
			status.SetCondition(av1.HelmResourceCondition{
				Type:         av1.ConditionFailed,
				Status:       av1.ConditionStatusTrue,
				Reason:       av1.ReasonInstallError,
				Message:      err.Error(),
				ResourceName: installedResource.GetName(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		// if spec.WatchHelmDependentResources && r.resourceWatchUpdater != nil {
		// 	if err := r.resourceWatchUpdater(installedResource); err != nil {
		// 		log.Error(err, "Failed to run update resource dependant resources")
		// 		return reconcile.Result{}, err
		// 	}
		// }

		log.Info("Installed resource", "resourceName", installedResource.GetName())
		r.recorder.Event(instance, corev1.EventTypeNormal, "Installed", fmt.Sprintf("Installed Resource %s", installedResource.GetName()))
		status.SetCondition(av1.HelmResourceCondition{
			Type:         av1.ConditionDeployed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonInstallSuccessful,
			Message:      "HarcodedMessage",
			ResourceName: installedResource.GetName(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	if manager.IsUpdateRequired() {
		previousResource, updatedResource, err := manager.UpdateResource(context.TODO())
		if previousResource != nil && updatedResource != nil {
			log.Info(previousResource.GetName(), updatedResource.GetName())
		}
		if err != nil {
			log.Error(err, "Failed to update resource")
			r.recorder.Event(instance, corev1.EventTypeWarning, "UpdateFailure", fmt.Sprintf("Updated Resource %s", updatedResource.GetName()))
			status.SetCondition(av1.HelmResourceCondition{
				Type:         av1.ConditionFailed,
				Status:       av1.ConditionStatusTrue,
				Reason:       av1.ReasonUpdateError,
				Message:      err.Error(),
				ResourceName: updatedResource.GetName(),
			})
			_ = r.updateResourceStatus(instance, status)
			return reconcile.Result{}, err
		}
		status.RemoveCondition(av1.ConditionFailed)

		// if spec.WatchHelmDependentResources && r.resourceWatchUpdater != nil {
		// 	if err := r.resourceWatchUpdater(updatedResource); err != nil {
		// 		log.Error(err, "Failed to run update resource dependant resources")
		// 		return reconcile.Result{}, err
		// 	}
		// }

		log.Info("Updated resource", "resourceName", updatedResource.GetName())
		r.recorder.Event(instance, corev1.EventTypeNormal, "Updated", fmt.Sprintf("Updated Resource %s", updatedResource.GetName()))
		status.SetCondition(av1.HelmResourceCondition{
			Type:         av1.ConditionDeployed,
			Status:       av1.ConditionStatusTrue,
			Reason:       av1.ReasonUpdateSuccessful,
			Message:      "HardcodedMessage",
			ResourceName: updatedResource.GetName(),
		})
		err = r.updateResourceStatus(instance, status)
		return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
	}

	// expectedResource, err := manager.ReconcileRelease(context.TODO())
	_, err = manager.ReconcileResource(context.TODO())
	if err != nil {
		log.Error(err, "Failed to reconcile resource")
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

	// if spec.WatchHelmDependentResources && r.resourceWatchUpdater != nil {
	// 	if err := r.resourceWatchUpdater(expectedResource); err != nil {
	// 		log.Error(err, "Failed to run update resource dependant resources")
	// 		return reconcile.Result{}, err
	// 	}
	// }

	log.Info("Reconciled resource")
	err = r.updateResourceStatus(instance, status)
	return reconcile.Result{RequeueAfter: r.reconcilePeriod}, err
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

	if log.Enabled() {
		fmt.Println(helper.PrettyPrint())
	}

	// JEB: Be sure to have update status subresources in the CRD.yaml
	// JEB: Look for kubebuilder subresources in the _types.go
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status. Ignoring")
		err = nil
	}

	return err
}
