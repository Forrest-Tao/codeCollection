package kubevirtservice

import (
	"context"
	virtcorev1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
)

type OSType string

const (
	Ubuntu  OSType = "Ubuntu"
	Windows OSType = "Windows"
)

// 重启VM
func RestartVM(ctx context.Context, virtClient kubecli.KubevirtClient, namespace, name string) error {
	return virtClient.VirtualMachine(namespace).Restart(ctx, name, &virtcorev1.RestartOptions{})
}
