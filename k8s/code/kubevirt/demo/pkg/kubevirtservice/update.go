package kubevirtservice

import (
	"context"
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/resource"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	virtcorev1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
)

func modifyVM(vm *virtcorev1.VirtualMachine) {
	vm.Spec.Template.Spec.Domain.Resources.Requests["cpu"] = resource.MustParse("3")      // 设置 CPU 请求为 2
	vm.Spec.Template.Spec.Domain.Resources.Requests["memory"] = resource.MustParse("2Gi") // 设置内存请求为 2Gi
}

func patchVMLabel(ctx context.Context, virtClient kubecli.KubevirtClient, namespace, vmName, labelKey, labelValue string) error {
	patchData := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]string{
				labelKey: labelValue,
			},
		},
	}
	patchBytes, _ := json.Marshal(patchData)
	_, err := virtClient.VirtualMachine(namespace).Patch(ctx, vmName, types.MergePatchType, patchBytes, k8smetav1.PatchOptions{})
	return err
}
