package deployment

import (
	"context"
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"os"
	"strings"
	"time"
)

var ctx = context.Background()

func ReadDeploymentYaml(filename string) appsv1.Deployment {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	deployJson, err := yaml.ToJSON(data)
	if err != nil {
		panic(err)
	}
	var deploy appsv1.Deployment
	if json.Unmarshal(deployJson, &deploy) != nil {
		panic(err)
	}
	return deploy
}

func ApplyDeployment(clientset kubernetes.Clientset, deploy appsv1.Deployment) {
	var namespace string
	if deploy.Namespace == "" {
		namespace = "default"
	} else {
		namespace = deploy.Namespace
	}

	if _, err := clientset.AppsV1().Deployments(namespace).Get(ctx, deploy.Name, metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			panic(err)
		}
		_, err = clientset.AppsV1().Deployments(namespace).Create(ctx, &deploy, metav1.CreateOptions{})
	} else {
		_, err = clientset.AppsV1().Deployments(namespace).Update(ctx, &deploy, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Deployment applied")
}

func DeleteDeployment(clientset kubernetes.Clientset, deployment appsv1.Deployment) {
	var namespace string
	if deployment.Namespace == "" {
		namespace = "default"
	} else {
		namespace = deployment.Namespace
	}
	err := clientset.AppsV1().Deployments(namespace).Delete(ctx, deployment.Name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}

func GetDeployment(clientset kubernetes.Clientset, name, namespace string) appsv1.Deployment {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return appsv1.Deployment{}
	}
	return *deployment
}

func GetDeploymentCondition(status appsv1.DeploymentStatus, conditionType appsv1.DeploymentConditionType) *appsv1.DeploymentCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == conditionType {
			return &c
		}
	}
	return nil
}

func GetPodCondition(status corev1.PodStatus, conditionType corev1.PodConditionType) *corev1.PodCondition {
	for i := range status.Conditions {
		c := status.Conditions[i]
		if c.Type == conditionType {
			return &c
		}
	}
	return nil
}

func printDeploymentStatus(clientset kubernetes.Clientset, deployment appsv1.Deployment) {
	labelSelector := ""
	for key, value := range deployment.Spec.Selector.MatchLabels {
		labelSelector += labelSelector + key + "=" + value + ","
	}

	labelSelector = strings.TrimRight(labelSelector, ",")
	var (
		err           error
		name          = deployment.Name
		namespace     = "default"
		k8sDeployment *appsv1.Deployment
	)
	if deployment.Namespace != "" {
		namespace = deployment.Namespace
	}
	for {
		k8sDeployment, err = clientset.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		// 打印deployment状态
		fmt.Printf("-------------deployment status------------\n")
		fmt.Printf("deployment.name: %s\n", k8sDeployment.Name)
		fmt.Printf("deployment.generation: %d\n", k8sDeployment.Generation)
		fmt.Printf("deployment.status.observedGeneration: %d\n", k8sDeployment.Status.ObservedGeneration)
		fmt.Printf("deployment.spec.replicas: %d\n", *(k8sDeployment.Spec.Replicas))
		fmt.Printf("deployment.status.replicas: %d\n", k8sDeployment.Status.Replicas)
		fmt.Printf("deployment.status.updatedReplicas: %d\n", k8sDeployment.Status.UpdatedReplicas)
		fmt.Printf("deployment.status.readyReplicas: %d\n", k8sDeployment.Status.ReadyReplicas)
		fmt.Printf("deployment.status.unavailableReplicas: %d\n", k8sDeployment.Status.UnavailableReplicas)

		podList, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: labelSelector})
		if err != nil {
			panic(err)
		}
		for index, pod := range podList.Items {
			// 打印pod的状态
			fmt.Printf("-------------pod %d status------------\n", index)
			fmt.Printf("pod.name: %s\n", pod.Name)
			fmt.Printf("pod.status.phase: %s\n", pod.Status.Phase)
			for _, condition := range pod.Status.Conditions {
				fmt.Printf("condition.type: %s, condition.status: %s, conditon.reason: %s\n", condition.Type, condition.Status, condition.Reason)
			}

			for _, containerStatus := range pod.Status.ContainerStatuses {
				if containerStatus.State.Waiting != nil {
					fmt.Printf("containerStatus.state.waiting.reason: %s\n", containerStatus.State.Waiting.Reason)
				}
				if containerStatus.State.Running != nil {
					fmt.Printf("containerStatus.state.running.startedAt: %s\n", containerStatus.State.Running.StartedAt)
				}
			}
		}
		availableCondition := GetDeploymentCondition(k8sDeployment.Status, appsv1.DeploymentAvailable)
		progressingCondition := GetDeploymentCondition(k8sDeployment.Status, appsv1.DeploymentProgressing)
		if k8sDeployment.Status.UpdatedReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.Replicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.AvailableReplicas == *(k8sDeployment.Spec.Replicas) &&
			k8sDeployment.Status.ObservedGeneration >= k8sDeployment.Generation &&
			availableCondition.Status == "True" &&
			progressingCondition.Status == "True" {
			fmt.Printf("-------------deploy status------------\n")
			fmt.Println("success!")
		} else {
			fmt.Printf("-------------deploy status------------\n")
			fmt.Println("waiting...")
		}

		time.Sleep(10 * time.Second)
	}
}

// TODO
func Gray(clientset kubernetes.Clientset, deployment appsv1.Deployment) {

}

// TODO
func Rollback(client kubernetes.Clientset) {

}
