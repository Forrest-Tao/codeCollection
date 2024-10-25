package service

import (
	"context"
	"encoding/json"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"os"
	"time"
)

var ctx = context.Background()

func ReadServiceYaml(filename string) coreV1.Service {
	serviceYaml, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	serviceJson, err := yaml.ToJSON(serviceYaml)
	if err != nil {
		panic(err)
	}
	var service coreV1.Service
	if json.Unmarshal(serviceJson, &service) != nil {
		panic(err)
	}
	return service
}

func ApplyServcice(clientset *kubernetes.Clientset, new_service coreV1.Service) (err error) {
	//err！=nil，说明service不存在或查询失败
	if _, err = clientset.CoreV1().Services(new_service.Namespace).Get(ctx, new_service.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			//说明查询出错
			return
		}
		//err！=nil&&IsNotFound()
		//创建service
		_, err = clientset.CoreV1().Services(new_service.Namespace).Create(ctx, &new_service, metav1.CreateOptions{})
		return
	}
	//err==nil 说明service存在，则更新
	_, err = clientset.CoreV1().Services(new_service.Namespace).Update(ctx, &new_service, metav1.UpdateOptions{})
	return
}

// 通过labelselector查询service
func GetServiceByLabelSelector(clientset *kubernetes.Clientset, namespace string, labelSelector string) (serviceList *coreV1.ServiceList, err error) {
	serviceList, err = clientset.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	return serviceList, nil
}

// watch service based on watcher
func WatchService(clientset *kubernetes.Clientset, namespace string, labelSelector string) {
	watcher, err := clientset.CoreV1().Services(namespace).Watch(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		panic(err)
	}
	defer watcher.Stop()
	for event := range watcher.ResultChan() {
		switch event.Type {
		//	Added    EventType = "ADDED"
		//	Modified EventType = "MODIFIED"
		//	Deleted  EventType = "DELETED"
		//	Bookmark EventType = "BOOKMARK"
		//	Error    EventType = "ERROR"
		case watch.Added:
			fmt.Printf("Service added: %s\n", event.Object.(*coreV1.Service).Name)
		case watch.Modified:
			fmt.Printf("Service modified: %s\n", event.Object.(*coreV1.Service).Name)
		case watch.Deleted:
			fmt.Printf("Service deleted: %s\n", event.Object.(*coreV1.Service).Name)
		case watch.Bookmark:
			fmt.Printf("Bookmark: %s\n", event.Object)
		case watch.Error:
			fmt.Printf("Error: %s\n", event.Object)
		}
	}
}

// watch service based on informer
func WatchServiceWithInformer(clientset *kubernetes.Clientset, namespace string, labelSelector string) {
	stopCh := make(chan struct{})
	defer close(stopCh)

	factory := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := factory.Core().V1().Services().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Printf("Service added: %s\n", obj.(*coreV1.Service).Name)
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Printf("Service updated: %s\n", new.(*coreV1.Service).Name)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Printf("Service deleted: %s\n", obj.(*coreV1.Service).Name)
		}})

	go informer.Run(stopCh)

	go func() {

	}()

	<-stopCh
}
