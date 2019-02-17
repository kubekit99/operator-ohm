package openstackrestore

import (
	"context"

	lcmv1alpha1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstackhelm/v1alpha1"
	lcmutils "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/controller/utils"
	//corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_openstackrestore")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new OpenstackRestore Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileOpenstackRestore{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("openstackrestore-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource OpenstackRestore
	err = c.Watch(&source.Kind{Type: &lcmv1alpha1.OpenstackRestore{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Workflows and requeue the owner OpenstackRestore
	o := lcmutils.NewWorkflowGroupVersionKind()
	err = c.Watch(&source.Kind{Type: o},
		&handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &lcmv1alpha1.OpenstackRestore{},
		},
		// predicate.GenerationChangedPredicate{}
	)
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileOpenstackRestore{}

// ReconcileOpenstackRestore reconciles a OpenstackRestore object
type ReconcileOpenstackRestore struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a OpenstackRestore object and makes changes based on the state read
// and what is in the OpenstackRestore.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Workflow as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileOpenstackRestore) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling OpenstackRestore")

	// Fetch the OpenstackRestore instance
	instance := &lcmv1alpha1.OpenstackRestore{}
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

	// Set OpenstackRestore instance as the owner and controller
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
			return reconcile.Result{}, err
		}

		// Workflow created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Workflow already exists - don't requeue
	reqLogger.Info("Skip reconcile: Workflow already exists", "Workflow.Namespace", found.GetNamespace(), "Workflow.Name", found.GetName())
	return reconcile.Result{}, nil
}
