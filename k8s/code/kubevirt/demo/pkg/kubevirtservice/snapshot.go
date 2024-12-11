package kubevirtservice

import (
	"codeCollection/k8s/code/kubevirt/demo/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1beta1 "kubevirt.io/api/snapshot/v1beta1"
	"time"
)

/*
apiVersion: snapshot.kubevirt.io/v1beta1
kind: VirtualMachineSnapshot
metadata:
  name: snap-larry
spec:
  source:
    apiGroup: kubevirt.io
    kind: VirtualMachine
    name: larr
*/

func CreateSnapshot(namespace, name string) *v1beta1.VirtualMachineSnapshot {
	return &v1beta1.VirtualMachineSnapshot{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"kubevirt.io/storage-version": "v1beta1",
			},
		},
		Spec: v1beta1.VirtualMachineSnapshotSpec{
			Source: corev1.TypedLocalObjectReference{
				APIGroup: utils.To[string]("kubevirt.io"),
				Kind:     "VirtualMachine",
				Name:     name,
			},
		},
	}
}

func createRestore(namespace, name string) *v1beta1.VirtualMachineRestore {
	return &v1beta1.VirtualMachineRestore{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:            "",
			GenerateName:    "",
			Namespace:       "",
			SelfLink:        "",
			UID:             "",
			ResourceVersion: "",
			Generation:      0,
			CreationTimestamp: metav1.Time{
				Time: time.Time{},
			},
			DeletionGracePeriodSeconds: nil,
			Labels:                     nil,
			Annotations:                nil,
			OwnerReferences:            nil,
			Finalizers:                 nil,
			ManagedFields:              nil,
		},
		Spec: v1beta1.VirtualMachineRestoreSpec{
			Target: corev1.TypedLocalObjectReference{
				APIGroup: nil,
				Kind:     "",
				Name:     "",
			},
			VirtualMachineSnapshotName: "",
			Patches:                    nil,
		},
	}
}
