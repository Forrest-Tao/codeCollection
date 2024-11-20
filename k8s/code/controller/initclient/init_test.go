package initclient

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestInit(t *testing.T) {
	list, err := ClientSet.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range list.Items {
		fmt.Println(n.Name)
	}
}
