package armada

import (
	"context"

	armadav1alpha1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a new ArmadaManifest Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func AddArmadaManifestController(mgr manager.Manager) error {

	r := &ArmadaManifestReconciler{
		client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("amf-recorder"),
	}

	// Create a new controller
	c, err := controller.New("amf-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArmadaManifest
	err = c.Watch(&source.Kind{Type: &armadav1alpha1.ArmadaManifest{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ArmadaManifest
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &armadav1alpha1.ArmadaManifest{},
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
	client   client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a ArmadaManifest object and makes changes based on the state read
// and what is in the ArmadaManifest.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ArmadaManifestReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ArmadaManifest")

	// Fetch the ArmadaManifest instance
	instance := &armadav1alpha1.ArmadaManifest{}
	instance.SetNamespace(request.Namespace)
	instance.SetName(request.Name)
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

	// Define a new Pod object
	pod := newPodForCRToDelete1(instance)

	// Set ArmadaManifest instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

func (r ArmadaManifestReconciler) updateResource(o *armadav1alpha1.ArmadaManifest) error {
	return r.client.Update(context.TODO(), o)
}

func (r ArmadaManifestReconciler) updateResourceStatus(instance *armadav1alpha1.ArmadaManifest, status *armadav1alpha1.ArmadaManifestStatus) error {
	reqLogger := log.WithValues("ArmadaManifest.Namespace", instance.Namespace, "ArmadaManifest.Name", instance.Name)

	// JEB: This is already a reference to the object
	// instance.Status = status

	// JEB: Be sure to have update status subresources in the CRD.yaml
	err := r.client.Status().Update(context.TODO(), instance)
	if err != nil {
		reqLogger.Error(err, "Failure to update status")
	}

	return err
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCRToDelete1(cr *armadav1alpha1.ArmadaManifest) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
