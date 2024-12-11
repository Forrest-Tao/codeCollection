package kubevirtservice

import (
	"context"
	"fmt"
	"kubevirt.io/client-go/kubecli"
)

func getVMGuestInfo(ctx context.Context, virtClient kubecli.KubevirtClient, namespace string, name string) {
	info, err := virtClient.VirtualMachineInstance(namespace).GuestOsInfo(ctx, name)
	if err != nil {
		panic(err)
	}
	fmt.Println(info.String())
}
