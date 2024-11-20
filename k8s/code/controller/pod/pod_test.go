package pod

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"sort"
	"strings"
)

// 列出在ns下的带有filtter前缀的所有podName和一个map，其中map的key为podName，value为这个podName对应的pod ptr
func ListPod(client kubernetes.Interface, namespace string, filter string) ([]string, map[string]*v1.Pod, error) {
	names := make([]string, 0)
	m := make(map[string]*v1.Pod, 0)
	ctx := context.Background()

	labelSet := labels.SelectorFromSet(labels.Set{
		"app": "nginx",
	})
	listOpts := metav1.ListOptions{
		LabelSelector: labelSet.String(),
		FieldSelector: "status.phase=Running",
		Limit:         100,
	}

	pods, err := client.CoreV1().Pods(namespace).List(ctx, listOpts)
	if err != nil {
		return nil, nil, err
	}

	for _, pod := range pods.Items {
		obj := pod
		name := obj.Name
		if filter == "" || strings.Contains(name, filter) && pod.DeletionTimestamp == nil {
			names = append(names, name)
			m[name] = &obj
		}
	}
	sort.Strings(names)
	return names, m, nil
}

// getPodNames 获取ns下的带有filter前缀的podNames
func getPodNames(client kubernetes.Interface, ns string, filter string) ([]string, error) {
	names := make([]string, 0)
	list, err := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return names, err
	}
	for _, pod := range list.Items {
		name := pod.Name
		if filter == "" || strings.Contains(name, filter) {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return names, nil
}

// getPodRestarts 获取指定pod的重启次数
func getPodRestarts(pod *v1.Pod) int32 {
	var restarts int32
	statuses := pod.Status.ContainerStatuses
	if len(statuses) == 0 {
		return restarts
	}
	for _, status := range statuses {
		restarts += status.RestartCount
	}
	return restarts
}
