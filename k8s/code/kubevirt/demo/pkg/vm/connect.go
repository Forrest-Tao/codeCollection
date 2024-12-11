package vm

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"kubevirt.io/client-go/kubecli"
)

//TODO ssh

//TODO vnc

//TODO RDP

// vnc连接
func VNCVmi(ctx context.Context, virtClient kubecli.KubevirtClient, ns, name string) error {
	streamInterface, err := virtClient.VirtualMachineInstance(ns).VNC(name)
	if err != nil {
		return err
	}
	//思路
	//
	_ = streamInterface
	return nil
}

/*
apiVersion: v1
kind: Service
metadata:

	creationTimestamp: "2024-11-07T02:48:52Z"
	name: cirros-ssh
	namespace: kubevirt
	resourceVersion: "195721127"
	uid: 0f82295a-ca3a-4b2d-9bc3-c2edd965d42e

spec:

	clusterIP: 10.102.30.93
	clusterIPs:
	- 10.102.30.93
	externalTrafficPolicy: Cluster
	internalTrafficPolicy: Cluster
	ipFamilies:
	- IPv4
	ipFamilyPolicy: SingleStack
	ports:
	- nodePort: 31421
	  port: 20222
	  protocol: TCP
	  targetPort: 22
	selector:
	  kubevirt.io/domain: testvm
	  kubevirt.io/size: small
	sessionAffinity: None
	type: NodePort

status:

	loadBalancer: {}
*/
func ExposeVMi(virtClient kubecli.KubevirtClient, namespace string, virtName string) error {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      virtName + "-ssh",
			Namespace: namespace,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       20222,
					TargetPort: intstr.FromInt32(22),
				},
			},
			Selector: map[string]string{
				"kubevirt.io/domain": virtName,
				"kubevirt.io/size":   "small",
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	_, err := virtClient.CoreV1().Services(namespace).Create(context.Background(), svc, metav1.CreateOptions{})
	return err
}

func ExposeRDP(virtName string) error {

	return nil
}
