package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/api/resource"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	virtcorev1 "kubevirt.io/api/core/v1"
	"kubevirt.io/client-go/kubecli"
	"log"
)

var virtClient kubecli.KubevirtClient
var ctx = context.Background()
var err error

func init() {
	clientConfig := kubecli.DefaultClientConfig(&pflag.FlagSet{})
	virtClient, err = kubecli.GetKubevirtClientFromClientConfig(clientConfig)
	if err != nil {
		log.Fatalf("cannot obtain KubeVirt client: %v\n", err)
	}
}

func main() {
	// 定义 VirtualMachine 结构体
	//vm := &virtcorev1.VirtualMachine{
	//	ObjectMeta: k8smetav1.ObjectMeta{
	//		Name:      "cirros-vm",
	//		Namespace: "kubevirt",
	//	},
	//	Spec: virtcorev1.VirtualMachineSpec{
	//		Running: boolPtr(true),
	//		Template: &virtcorev1.VirtualMachineInstanceTemplateSpec{
	//			ObjectMeta: k8smetav1.ObjectMeta{
	//				Labels: map[string]string{
	//					"kubevirt.io/size":   "small",
	//					"kubevirt.io/domain": "testvm",
	//				},
	//			},
	//			Spec: virtcorev1.VirtualMachineInstanceSpec{
	//				Domain: virtcorev1.DomainSpec{
	//					Devices: virtcorev1.Devices{
	//						Disks: []virtcorev1.Disk{
	//							{
	//								Name: "containerdisk",
	//								DiskDevice: virtcorev1.DiskDevice{
	//									Disk: &virtcorev1.DiskTarget{
	//										Bus: "virtio",
	//									},
	//								},
	//							},
	//						},
	//						Interfaces: []virtcorev1.Interface{
	//							{
	//								Name: "default",
	//								InterfaceBindingMethod: virtcorev1.InterfaceBindingMethod{
	//									Masquerade: &virtcorev1.InterfaceMasquerade{},
	//								},
	//							},
	//						},
	//					},
	//					Resources: virtcorev1.ResourceRequirements{
	//						Requests: k8scorev1.ResourceList{
	//							k8scorev1.ResourceMemory: resource.MustParse("64M"),
	//						},
	//					},
	//				},
	//				Networks: []virtcorev1.Network{
	//					{
	//						Name: "default",
	//						NetworkSource: virtcorev1.NetworkSource{
	//							Pod: &virtcorev1.PodNetwork{},
	//						},
	//					},
	//				},
	//				Volumes: []virtcorev1.Volume{
	//					{
	//						Name: "containerdisk",
	//						VolumeSource: virtcorev1.VolumeSource{
	//							ContainerDisk: &virtcorev1.ContainerDiskSource{
	//								Image: "quay.io/kubevirt/cirros-container-volume-demo",
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//// 创建虚拟机
	//_, err = virtClient.VirtualMachine("kubevirt").Create(context.Background(), vm, k8smetav1.CreateOptions{})
	//if err != nil {
	//	log.Fatalf("cannot create KubeVirt VM: %v\n", err)
	//} else {
	//	fmt.Println("Virtual Machine created successfully.")
	//}
	//
	vm, err := virtClient.VirtualMachine("kubevirt").Get(ctx, "cirros-vm", k8smetav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	modifyVM(vm)
	if err = updateVM(vm, "kubevirt"); err != nil {
		log.Fatalf("cannot update VM: %v\n", err)
	}
	fmt.Println("update successful")
	if err = restartVMi(vm, "kubevirt"); err != nil {
		log.Fatalf("cannot restart VMi: %v\n", err)
	}
	fmt.Println("restart successful")
}

func boolPtr(b bool) *bool {
	return &b
}

func updateVM(vm *virtcorev1.VirtualMachine, namespace string) error {
	_, err = virtClient.VirtualMachine(namespace).Update(ctx, vm, k8smetav1.UpdateOptions{})
	return err
}

func getVM(namespace, name string) (*virtcorev1.VirtualMachine, error) {
	return virtClient.VirtualMachine(namespace).Get(ctx, name, k8smetav1.GetOptions{})
}

func restartVMi(vm *virtcorev1.VirtualMachine, namespace string) error {
	return virtClient.VirtualMachine(namespace).Restart(ctx, "cirros-vm", &virtcorev1.RestartOptions{})
}

func modifyVM(vm *virtcorev1.VirtualMachine) {
	vm.Spec.Template.Spec.Domain.Resources.Requests["cpu"] = resource.MustParse("3")      // 设置 CPU 请求为 2
	vm.Spec.Template.Spec.Domain.Resources.Requests["memory"] = resource.MustParse("2Gi") // 设置内存请求为 2Gi
}

func patchVMLabel(namespace, vmName, labelKey, labelValue string) error {
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
