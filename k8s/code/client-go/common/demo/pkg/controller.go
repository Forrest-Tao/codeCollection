package pkg

import (
	v12 "k8s.io/api/networking/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informer "k8s.io/client-go/informers/core/v1"
	netiIformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
)

const (
	workNum  = 5
	maxRetry = 10
)

type Controller struct {
	clientset     kubernetes.Interface
	ingressLister v1.IngressLister
	serviceLister corev1.ServiceLister
	queue         workqueue.TypedRateLimitingInterface[string]
}

func NewController(clientset kubernetes.Interface, serviceInformer informer.ServiceInformer, ingressInformer netiIformer.IngressInformer) Controller {
	c := Controller{
		clientset:     clientset,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue: workqueue.NewTypedRateLimitingQueueWithConfig(workqueue.DefaultTypedControllerRateLimiter[string](), workqueue.TypedRateLimitingQueueConfig[string]{
			Name: "ingressManager",
		}),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})
	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerDetailedFuncs{
		DeleteFunc: c.deleteIngress,
	})
	for i := 0; i < ; i++ {
		
	}
	return c
}

func (c Controller) enqueue(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c Controller) addService(obj interface{}) {
	c.enqueue(obj)
}

func (c Controller) updateService(oldObj interface{}, newObj interface{}) {
	//todo: compare annotations
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}

func (c Controller) deleteIngress(obj interface{}) {
	ingress := obj.(*v12.Ingress)
	owerReference := v13.GetControllerOf(ingress)
	if owerReference == nil {
		return
	}
	if owerReference.Kind != "Service" {
		return
	}
	c.enqueue(obj)
}

func (c Controller) Run(stopCh chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.worker, 0, stopCh)
	}
}

func (c *Controller) worker() {
	for c.processNextWorkItem() {

	}
}
func (c *Controller) processNextWorkItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)
	key := item
	err := c.syncService(key)
	if err == nil {
		c.handleError(key, err)
	}
	return true
}

func (c Controller) handleError(key string, err error) {
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c Controller) syncService(key string) interface{} {

}
