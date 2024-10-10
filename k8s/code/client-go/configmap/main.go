package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/Zhuanz/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func main() {
	clientset, err := createClient()
	if err != nil {
		panic(err)
	}
	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-demo",
		},
		Data: map[string]string{
			"key":        "value",
			"key2":       "value2",
			"file.txt":   "name=John\nage=30",
			"mysql.conf": "color.good=blue\ncolor.good=red\ncolor.bad=ok",
		},
	}

	// 创建 ConfigMap
	_, err = clientset.CoreV1().ConfigMaps("default").Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("ConfigMap created successfully.")

	// 获取 ConfigMap
	cm, err := clientset.CoreV1().ConfigMaps("default").Get(context.TODO(), "configmap-demo", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("ConfigMap Data: %v\n", cm.Data)
	//
	//// 删除 ConfigMap
	//err = clientset.CoreV1().ConfigMaps("default").Delete(context.TODO(), "configmap-demo", metav1.DeleteOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println("ConfigMap deleted successfully.")
}
