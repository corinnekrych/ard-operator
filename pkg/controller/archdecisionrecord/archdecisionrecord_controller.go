package archdecisionrecord

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"strconv"

	corinnekrychv1alpha1 "github.com/corinnekrych/adr-operator/pkg/apis/corinnekrych/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	//"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	// api clientsets
	//appsschema "github.com/openshift/client-go/apps/clientset/versioned/scheme"
	//appsclientset "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"
	//buildschema "github.com/openshift/client-go/build/clientset/versioned/scheme"
	//buildclientset "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	//imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	imagev1 "github.com/openshift/api/image/v1"
	//projectclientset "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	//routeclientset "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	//userclientset "github.com/openshift/client-go/user/clientset/versioned/typed/user/v1"
)

var log = logf.Log.WithName("controller_archdecisionrecord")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ArchDecisionRecord Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileArchDecisionRecord{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("archdecisionrecord-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ArchDecisionRecord
	err = c.Watch(&source.Kind{Type: &corinnekrychv1alpha1.ArchDecisionRecord{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ArchDecisionRecord
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &corinnekrychv1alpha1.ArchDecisionRecord{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileArchDecisionRecord{}

// ReconcileArchDecisionRecord reconciles a ArchDecisionRecord object
type ReconcileArchDecisionRecord struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a ArchDecisionRecord object and makes changes based on the state read
// and what is in the ArchDecisionRecord.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileArchDecisionRecord) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ArchDecisionRecord")

	// Fetch the ArchDecisionRecord instance
	instance := &corinnekrychv1alpha1.ArchDecisionRecord{}
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
	log.Info("============================================================")
	log.Info(fmt.Sprintf("***** Reconciling ArchDecisionRecord %s, namespace %s", request.Name, request.Namespace))
	log.Info(fmt.Sprintf("** Steps of the component : %s", instance.Status.Steps))
	log.Info(fmt.Sprintf("** Creation time : %s", instance.ObjectMeta.CreationTimestamp))
	log.Info(fmt.Sprintf("** Resource version : %s", instance.ObjectMeta.ResourceVersion))
	log.Info(fmt.Sprintf("** Generation version : %s", strconv.FormatInt(instance.ObjectMeta.Generation, 10)))
	log.Info(fmt.Sprintf("** Deletion time : %s", instance.ObjectMeta.DeletionTimestamp))
	log.Info("============================================================")

	// Define a new Image object
	image, err := r.newImageStream(instance.Spec.Image, instance)
	if err != nil {
		log.Error(err, "Creating image stream fails")
		return reconcile.Result{}, err
	}
	// Set ArchDecisionRecord instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, image, r.scheme); err != nil {
		log.Error(err, "Setting owner reference fails")
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	//found := &corev1.Pod{}
	//err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	//if err != nil && errors.IsNotFound(err) {
	//	reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
	//	err = r.client.Create(context.TODO(), pod)
	//	if err != nil {
	//		return reconcile.Result{}, err
	//	}
	//
	//	// Pod created successfully - don't requeue
	//	return reconcile.Result{}, nil
	//} else if err != nil {
	//	return reconcile.Result{}, err
	//}

	// Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}


func (r *ReconcileArchDecisionRecord) newImageStream(imageName string, instance *corinnekrychv1alpha1.ArchDecisionRecord) (*imagev1.ImageStream, error) {

	imageStream := newImageStream(instance)
	found := &imagev1.ImageStream{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: imageStream.Name, Namespace: imageStream.Namespace}, found)
	if err != nil {
		imageStream.Namespace = "openshift"
		errOS := r.client.Get(context.TODO(), types.NamespacedName{Name: imageStream.Name, Namespace: imageStream.Namespace}, found)
		if err != nil {
			return nil, errOS
		}
	}
	return found, nil

	// Without using split client
	//namespaceImageStream, err1 = r.imageClient.ImageStreams(namespace).Get(imageName, metav1.GetOptions{})
	//if err1 != nil {
	//	openshiftImageStream, err2 = r.imageClient.ImageStreams(namespace).Get(imageName, metav1.GetOptions{})
	//	if err2 != nil {
	//		return nil, err2
	//	} else {
	//		return openshiftImageStream, nil
	//	}
	//} else {
	//	return namespaceImageStream, nil
	//}

	//return imageStream, nil
}

func newImageStream(instance *corinnekrychv1alpha1.ArchDecisionRecord) *imagev1.ImageStream {
	labels := map[string]string{
		"app": instance.Name,
	}
	return &imagev1.ImageStream{ObjectMeta:metav1.ObjectMeta{
		Name: instance.Spec.Image,
		Namespace: instance.Namespace,
		Labels:    labels,
	}}
}
// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *corinnekrychv1alpha1.ArchDecisionRecord) *corev1.Pod {
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
