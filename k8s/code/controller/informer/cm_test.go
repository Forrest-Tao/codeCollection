package informer

import (
	"fmt"
	"forrest/codeCollection/k8s/code/controller/initclient"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"testing"
)

func TestCm(t *testing.T) {
	t.Log("Testing CM")
	client := initclient.GetClient().Client
	factory := informers.NewSharedInformerFactory(client, 0)
	cmInformer := factory.Core().V1().ConfigMaps().Informer()
	_, err := cmInformer.AddEventHandler(&cmHandler{})
	if err != nil {
		t.Error(err)
	}

	//cmInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc:    nil,
	//	UpdateFunc: nil,
	//	DeleteFunc: nil,
	//})

	//start 多个informer同时start，非阻塞，仅发出信号
	//factory.Start(wait.NeverStop)
	//factory.WaitForCacheSync(wait.NeverStop)
	//select {
	//
	//}

	//run,某个informer run起来，阻塞，正式监听对象的变化
	stopCh := make(chan struct{})
	defer close(stopCh)
	cmInformer.Run(stopCh)
}

type cmHandler struct {
}

func (c cmHandler) OnAdd(obj interface{}, isInInitialList bool) {
	fmt.Println("onAdd", isInInitialList, obj.(*v1.ConfigMap).Name)
}

func (c cmHandler) OnUpdate(oldObj, newObj interface{}) {
	fmt.Println("onUpdate", oldObj.(*v1.ConfigMap).Name, newObj.(*v1.ConfigMap).Name)
}

func (c cmHandler) OnDelete(obj interface{}) {
	fmt.Println("onDelete", obj.(*v1.ConfigMap).Name)
}
