package watch

import (
	"context"
	"fmt"
	"forrest/codeCollection/k8s/code/controller/initclient"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"testing"
	"time"
)

func TestWatchPods(t *testing.T) {
	client := initclient.ClientSet.Client

	watch, err := client.CoreV1().Pods("default").
		Watch(context.TODO(), metav1.ListOptions{
			LabelSelector: "app==testApp", // 只监控带有特定标签的Pod
		})
	if err != nil {
		fmt.Println("Error starting watch:", err)
		return
	}
	defer watch.Stop()

	// 启动协程来处理事件
	go func() {
		for event := range watch.ResultChan() {
			pod, ok := event.Object.(*corev1.Pod)
			if !ok {
				fmt.Println("unexpected type")
				continue
			}

			switch event.Type {
			case "ADDED":
				fmt.Printf("Pod added: %s/%s\n", pod.Namespace, pod.Name)
			case "MODIFIED":
				fmt.Printf("Pod modified: %s/%s\n", pod.Namespace, pod.Name)
			case "DELETED":
				fmt.Printf("Pod deleted: %s/%s\n", pod.Namespace, pod.Name)
			default:
				fmt.Printf("Unknown event: %v\n", event.Type)
			}
		}
	}()

	pod1 := createPod(client, "test-pod-1")
	pod2 := createPod(client, "test-pod-2")

	deletePod(client, pod1)
	deletePod(client, pod2)

	time.Sleep(10 * time.Second)
}

func createPod(client kubernetes.Interface, name string) *corev1.Pod {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
			Labels: map[string]string{
				"app": "testApp",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "busybox",
					Image: "busybox",
					Command: []string{
						"sh",
						"-c",
						"sleep 3600",
					},
				},
			},
		},
	}

	pod, err := client.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("create pod err:", err)
		return nil
	}
	return pod
}

func deletePod(client kubernetes.Interface, pod *corev1.Pod) {
	err := client.CoreV1().Pods(pod.Namespace).Delete(context.Background(), pod.Name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("delete pod err:", err)
		return
	}
}
