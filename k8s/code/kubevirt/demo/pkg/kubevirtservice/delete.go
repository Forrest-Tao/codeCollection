package kubevirtservice

import (
	"context"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubevirt.io/client-go/kubecli"
)

func DeleteVM(ctx context.Context, virtClient kubecli.KubevirtClient, namespace, name string) error {
	return virtClient.VirtualMachine(namespace).Delete(ctx, name, k8smetav1.DeleteOptions{})
}
