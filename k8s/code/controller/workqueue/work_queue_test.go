package workqueue

import (
	"context"
	"fmt"
	"forrest/codeCollection/k8s/code/controller/initclient"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"log"
	"testing"
)

var (
	workerNum int = 3
)

func getQueue[T comparable]() workqueue.TypedRateLimitingInterface[T] {
	return workqueue.NewTypedRateLimitingQueue[T](workqueue.DefaultTypedControllerRateLimiter[T]())
}

func TestWorkQueue(t *testing.T) {
	client := initclient.ClientSet.Client
	queue := getQueue[string]()
	defer queue.ShutDown()

	factory := informers.NewSharedInformerFactory(client, 0)
	cmInformer := factory.Core().V1().ConfigMaps().Informer()
	cmLister := factory.Core().V1().ConfigMaps().Lister()

	cmInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("new event : ADD ", key)
				//just push into queue
				queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err == nil {
				fmt.Println("new event : UPDATE ", key)
				//just push into queue
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				fmt.Println("delete event : DELETE ", key)
				//just push into queue
				queue.Add(key)
			}
		},
	})

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	factory.Start(ctx.Done())

	//wait for resource sync
	for gvr, ok := range factory.WaitForCacheSync(ctx.Done()) {
		if !ok {
			log.Fatal(fmt.Sprintf("Failed to sync cache for resource %v", gvr))
		}
		log.Println(fmt.Sprintf("cache sync finished for resource %v", gvr))
	}

	//start workerNum goroutine to handle
	for i := 0; i < workerNum; i++ {
		go func(n int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("worker exit")
					return
				default:
				}
				key, quit := queue.Get()
				if quit {
					fmt.Printf("Work queue has been shut down! Worker %d exiting...\n", n)
					return
				}
				fmt.Printf("Worker %d is about to start processing item %s.\n", n, key)

				func() {
					defer queue.Done(key)

					namespace, name, err := cache.SplitMetaNamespaceKey(key)
					if err != nil {
						fmt.Printf("Worker %d failed to split key %s: %v\n", n, key, err)
						return
					}

					cm, err := cmLister.ConfigMaps(namespace).Get(name)

					//if the key has been handled successfully -forget it
					if err == nil {
						fmt.Printf("Worker %d successfully processed ConfigMap: %s/%s with data: %v\n", n, namespace, name, cm.Data)
						queue.Forget(key)
						return
					}

					//retry more than 3,just forget it s
					if queue.NumRequeues(key) >= 3 {
						fmt.Printf("Worker %d failed to process ConfigMap after 3 retries: %v\n", n, err)
						queue.Forget(key)
						return
					}
					//failed to handle,retry
					fmt.Printf("Worker %d failed to process ConfigMap after 3 retries: %v\n", n, err)
					queue.AddRateLimited(key)

				}()
			}
		}(i)
	}
	select {}
}
