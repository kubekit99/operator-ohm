package openstackrollback

import (
	"context"
	"fmt"

	lcmv1alpha1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstackhelm/v1alpha1"
	lcmutils "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/controller/utils"
	//corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_openstackrollback")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new OpenstackRollback Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileOpenstackRollback{client: mgr.GetClient(), scheme: mgr.GetScheme(), recorder: mgr.GetRecorder("openstackrollback-recorder")}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("openstackrollback-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource OpenstackRollback
	err = c.Watch(&source.Kind{Type: &lcmv1alpha1.OpenstackRollback{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Workflows and requeue the owner OpenstackRollback
	o := lcmutils.NewWorkflowGroupVersionKind()
	err = c.Watch(&source.Kind{Type: o},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &lcmv1alpha1.OpenstackRollback{},
		},
		// predicate.GenerationChangedPredicate{}
	)
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileOpenstackRollback{}

// ReconcileOpenstackRollback reconciles a OpenstackRollback object
type ReconcileOpenstackRollback struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client   client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a OpenstackRollback object and makes changes based on the state read
// and what is in the OpenstackRollback.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Workflow as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileOpenstackRollback) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling OpenstackRollback")

	// Fetch the OpenstackRollback instance
	instance := &lcmv1alpha1.OpenstackRollback{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Workflow object
	wf := lcmutils.NewWorkflowForCR(instance.Name, instance.Namespace)

	// name of your custom finalizer
	myFinalizerName := "finalizer.openstackrollback"

	if instance.ObjectMeta.DeletionTimestamp.IsZero() {

		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object.
		if !lcmutils.FinalizerContainsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			reqLogger.Info("Handling OpenstackRollback creation/update. Adding Finalizer")
			r.recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("Adding Finalizier"))
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.client.Update(context.Background(), instance); err != nil {
				return reconcile.Result{}, err
			}
		} else {
			reqLogger.Info("Handling OpenstackRollback creation/update")

		}
	} else {

		// The object is being deleted
		if lcmutils.FinalizerContainsString(instance.ObjectMeta.Finalizers, myFinalizerName) {
			reqLogger.Info("Handling OpenstackRollback deletion step 1")
			reqLogger.Info("Deleting Workflow", "Workflow.Namespace", wf.GetNamespace(), "Workflow.Name", wf.GetName(), "Worflow.Kind", wf.GetKind())

			// our finalizer is present, so lets handle our external dependency
			if err := r.deleteWorkflow(wf); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				// r.recorder.Event(instance, "Normal", "Failure", fmt.Sprintf("Deleting worfklow %s/%s", wf.GetNamespace(), wf.GetName()))
				return reconcile.Result{}, err
			} else {
				// r.recorder.Event(instance, "Normal", "Deleted", fmt.Sprintf("Deleting worfklow %s/%s", wf.GetNamespace(), wf.GetName()))
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = lcmutils.FinalizerRemoveString(instance.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.client.Update(context.Background(), instance); err != nil {
				return reconcile.Result{}, err
			}
		} else {
			reqLogger.Info("Handling OpenstackRollback deletion step 2")
		}

		// Our finalizer has finished, so the reconciler can do nothing.
		return reconcile.Result{}, nil
	}

	// Set OpenstackRollback instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, wf, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Workflow already exists
	found := lcmutils.NewWorkflowGroupVersionKind()
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: wf.GetName(), Namespace: wf.GetNamespace()}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Workflow", "Workflow.Namespace", wf.GetNamespace(), "Workflow.Name", wf.GetName(), "Worflow.Kind", wf.GetKind())

		err = r.client.Create(context.TODO(), wf)
		if err != nil {
			r.recorder.Event(instance, "Normal", "Failure", fmt.Sprintf("Creating worfklow %s/%s", wf.GetNamespace(), wf.GetName()))
			return reconcile.Result{}, err
		} else {
			r.recorder.Event(instance, "Normal", "Created", fmt.Sprintf("Creating worfklow %s/%s", wf.GetNamespace(), wf.GetName()))
		}

		// Workflow created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		r.recorder.Event(instance, "Normal", "Failure", fmt.Sprintf("Retrieving worfklow %s/%s", wf.GetNamespace(), wf.GetName()))
		return reconcile.Result{}, err
	}

	// Workflow already exists - don't requeue
	return reconcile.Result{}, nil
}

func (r *ReconcileOpenstackRollback) deleteWorkflow(wf *unstructured.Unstructured) error {
	found := lcmutils.NewWorkflowGroupVersionKind()
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: wf.GetName(), Namespace: wf.GetNamespace()}, found)
	if err != nil && errors.IsNotFound(err) {
		// Workflow was already deleted - don't requeue
		return nil
	} else if err != nil {
		return err
	}

	err = r.client.Delete(context.TODO(), found)
	if err != nil {
		return err
	}

	return nil
}
