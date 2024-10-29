package main

import (
	"context"
	"fmt"
	"github.com/spf13/pflag"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	corev1 "kubevirt.io/api/core/v1"
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
	virtClient.VirtualMachine("kubevirt").Create()
	vm, err := virtClient.VirtualMachine("kubevirt").Get(ctx, "cirros-vm-vm", k8smetav1.GetOptions{})
	fmt.Println(vm.Name)
	var newvm = new(corev1.VirtualMachine)
	vm.DeepCopyInto(newvm)
	vm.Name = "new-vm"
	err = virtClient.VirtualMachine("kubevirt").Delete(ctx, "cirros-vm", k8smetav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("cannot delete KubeVirt VM: %v\n", err)
	}
	_, err = virtClient.VirtualMachine("kubevirt").Create(ctx, newvm, k8smetav1.CreateOptions{})
	if err != nil {
		log.Fatalf("cannot create KubeVirt VM: %v\n", err)
	}
}
