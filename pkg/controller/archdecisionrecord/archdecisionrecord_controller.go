package archdecisionrecord

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"math/rand"
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

	appsv1 "github.com/openshift/api/apps/v1"
	buildv1 "github.com/openshift/api/build/v1"
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

	// 1) Does the image stream exist on current namespace or "openshift" namespace?
	_, errStream := r.GetImageStream(instance.Spec.Image, instance)
	if errStream != nil && errors.IsNotFound(errStream) {
		log.Error(err, ":::::::: No image stream found ::::::::")
		return reconcile.Result{}, err
	}
	// 2) Create an empty image name "nodejs-generated-xxxx"
	rand := "xxxx"//randomString(5)
	newImage := newImageStreamObjectMeta(instance.Namespace, instance.Name, fmt.Sprintf("%s-%s", instance.Spec.Image + "-generated", rand))
	err = r.client.Create(context.TODO(), newImage)
	if err != nil {
		log.Error(err, ":::::::: Creating new image fails ::::::::")
		return reconcile.Result{}, err
	}
	log.Info(":::::::: Image stream created ::::::::")
	// Set ArchDecisionRecord instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, newImage, r.scheme); err != nil {
		log.Error(err, ":::::::: Setting owner reference fails ::::::::")
		return reconcile.Result{}, err
	}

	// 3) Create build config with s2i
	imageNamespace := "openshift" // todo fix that
	bc := generateBuildConfig(instance.Name, instance.Spec.Source, "", instance.Spec.Image + ":latest", imageNamespace)
	err = r.client.Create(context.TODO(), &bc)
	if err != nil {
		log.Error(err, ":::::::: Creating build config fails ::::::::")
		return reconcile.Result{}, err
	}
	log.Info(":::::::: Build config created ::::::::")

	// 4) Create deployment config
    dc := generateGitDeploymentConfig(getMetaObj(instance.Name, imageNamespace), "nodejs-generated-xxxx:latest")
	err = r.client.Create(context.TODO(), &dc)
	if err != nil {
		log.Error(err, ":::::::: Creating deployment config fails ::::::::")
		return reconcile.Result{}, err
	}
	log.Info(":::::::: Build deployment created ::::::::")

	// 5) Start build
	retrievedBC := buildv1.BuildConfig{ObjectMeta: getMetaObj(instance.Name, instance.Namespace)} // todo labels selector?
	errFetchBC := r.client.Get(context.TODO(), request.NamespacedName, &retrievedBC)
	if errFetchBC != nil {
		log.Error(err, ":::::::: Fetching build config (to instantiate build) fails ::::::::")
		return reconcile.Result{}, err
	}
	// todo use k8s api to instantiate build
	log.Info(":::::::: Fetching build config succeed ::::::::")
	//retrievedBC
	// Pod already exists - don't requeue
	//reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}

func (r *ReconcileArchDecisionRecord) GetImageStream(imageName string, instance *corinnekrychv1alpha1.ArchDecisionRecord) (*imagev1.ImageStream, error) {
	imageStream := newImageStreamObjectMeta(instance.Namespace, instance.Name, instance.Spec.Image)
	found := &imagev1.ImageStream{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: imageStream.Name, Namespace: imageStream.Namespace}, found)
	if err != nil {
		log.Error(err, fmt.Sprintf("::Searching in namespace %s imagestream %s fails", imageStream.Namespace, imageStream.Name))
		imageStream.Namespace = "openshift"
		errOS := r.client.Get(context.TODO(), types.NamespacedName{Name: imageStream.Name, Namespace: imageStream.Namespace}, found)
		//errOS := r.client.List(context.TODO(), &client.ListOptions{}, found)
		if err != nil {
			log.Error(err, fmt.Sprintf("::listing in namespace %s imagestream %s fails", imageStream.Namespace, imageStream.Name))
			return nil, errOS
		}

	}
	log.Info(fmt.Sprintf("::::::::::::::::::::::::::: Found imageStream %s ::::::::::::::::::", found.ObjectMeta.ResourceVersion))
	return found, nil
}

func newImageStreamObjectMeta(namespace string, name string, imageName string) *imagev1.ImageStream {
	labels := map[string]string{
		"app": name,
	}
	return &imagev1.ImageStream{ObjectMeta:metav1.ObjectMeta{
		Name: imageName,
		Namespace: namespace,
		Labels:    labels,
	}}
}

func getMetaObj(name string, imageNamespace string) metav1.ObjectMeta {
	labels := map[string]string{
		"app": name,
	}
	return metav1.ObjectMeta{Name: name, Namespace:imageNamespace, Labels: labels}
}
// generateBuildConfig creates a BuildConfig for Git URL's being passed into Odo
func generateBuildConfig(name string, gitURL, gitRef, imageName, imageNamespace string) buildv1.BuildConfig {
	buildSource := buildv1.BuildSource{
		Git: &buildv1.GitBuildSource{
			URI: gitURL,
			Ref: gitRef,
		},
		Type: buildv1.BuildSourceGit,
	}

	return buildv1.BuildConfig{
		ObjectMeta: getMetaObj(name, imageNamespace),
		Spec: buildv1.BuildConfigSpec{
			CommonSpec: buildv1.CommonSpec{
				Output: buildv1.BuildOutput{
					To: &corev1.ObjectReference{
						Kind: "ImageStreamTag",
						Name: name + ":latest",
					},
				},
				Source: buildSource,
				Strategy: buildv1.BuildStrategy{
					SourceStrategy: &buildv1.SourceBuildStrategy{
						From: corev1.ObjectReference{
							Kind:      "ImageStreamTag",
							Name:      imageName,
							Namespace: imageNamespace,
						},
					},
				},
			},
		},
	}
}

func generateGitDeploymentConfig(commonObjectMeta metav1.ObjectMeta, image string) appsv1.DeploymentConfig {
	ports := make([]corev1.ContainerPort, 0)
	ports = append(ports, corev1.ContainerPort{Name: "8080-tcp", ContainerPort:8080})

	dc := appsv1.DeploymentConfig{
		ObjectMeta: commonObjectMeta,
		Spec: appsv1.DeploymentConfigSpec{
			Replicas: 1,
			Strategy: appsv1.DeploymentStrategy{
				Type: appsv1.DeploymentStrategyTypeRecreate,
			},
			Selector: map[string]string{
				"deploymentconfig": commonObjectMeta.Name,
			},
			Template: &corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"deploymentconfig": commonObjectMeta.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Image: image,
							Name:  commonObjectMeta.Name,
							Ports: ports,
							Env:   nil, // todo
						},
					},
				},
			},
			Triggers: []appsv1.DeploymentTriggerPolicy{
				{
					Type: "ConfigChange",
				},
				{
					Type: "ImageChange",
					ImageChangeParams: &appsv1.DeploymentTriggerImageChangeParams{
						Automatic: true,
						ContainerNames: []string{
							commonObjectMeta.Name,
						},
						From: corev1.ObjectReference{
							Kind: "ImageStreamTag",
							Name: image,
						},
					},
				},
			},
		},
	}
	// todo
	//containerIndex := -1
	//if resourceRequirements != nil {
	//	for index, container := range dc.Spec.Template.Spec.Containers {
	//		if container.Name == commonObjectMeta.Name {
	//			containerIndex = index
	//			break
	//		}
	//	}
	//	if containerIndex != -1 {
	//		dc.Spec.Template.Spec.Containers[containerIndex].Resources = *resourceRequirements
	//	}
	//}
	return dc
}