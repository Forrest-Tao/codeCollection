package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var ctx = context.Background()

func GetPodsByNamespace(clientset kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	return clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
}

func GetContainerLog(clientset kubernetes.Clientset, namespace string, podName string, container string) {
	res := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{Container: container}).Do(ctx)
	if res.Error() != nil {
		fmt.Println(res.Error())
		return
	}
	logs, err := res.Raw()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(logs))
}

//exec

//logs

//events
