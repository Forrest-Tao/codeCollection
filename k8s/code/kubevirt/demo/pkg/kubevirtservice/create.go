package kubevirtservice

import (
	"codeCollection/k8s/code/kubevirt/demo/utils"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	virtcorev1 "kubevirt.io/api/core/v1"
	"kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

//服务的暴露
//windows  RDP服务暴露  -> service ingress
//ubuntu   ssh服务暴露  -> service ingress

func createDataVolume(sys, data string) []virtcorev1.DataVolumeTemplateSpec {
	return []virtcorev1.DataVolumeTemplateSpec{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: sys,
			},
			Spec: v1beta1.DataVolumeSpec{
				Source: &v1beta1.DataVolumeSource{
					HTTP: &v1beta1.DataVolumeSourceHTTP{
						URL:                "",
						SecretRef:          "",
						CertConfigMap:      "",
						ExtraHeaders:       nil,
						SecretExtraHeaders: nil,
					},
					S3: &v1beta1.DataVolumeSourceS3{
						URL:           "",
						SecretRef:     "",
						CertConfigMap: "",
					},
					GCS: &v1beta1.DataVolumeSourceGCS{
						URL:       "",
						SecretRef: "",
					},
					Registry: &v1beta1.DataVolumeSourceRegistry{
						URL: utils.To[string]("docker://quay.io/kubevirt/cirros-container-disk-demo"),
					},
					PVC: &v1beta1.DataVolumeSourcePVC{
						Namespace: "",
						Name:      "",
					},
					Upload: &v1beta1.DataVolumeSourceUpload{},
					Blank:  &v1beta1.DataVolumeBlankImage{},
					Imageio: &v1beta1.DataVolumeSourceImageIO{
						URL:           "",
						DiskID:        "",
						SecretRef:     "",
						CertConfigMap: "",
					},
					VDDK: &v1beta1.DataVolumeSourceVDDK{
						URL:          "",
						UUID:         "",
						BackingFile:  "",
						Thumbprint:   "",
						SecretRef:    "",
						InitImageURL: "",
					},
					Snapshot: &v1beta1.DataVolumeSourceSnapshot{
						Namespace: "",
						Name:      "",
					},
				},
				Storage: &v1beta1.StorageSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{
						corev1.ReadWriteOnce,
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: resource.MustParse("2Gi"),
						},
					},
					StorageClassName: utils.To[string]("openebs-localpv"),
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: data,
			},
			Spec: v1beta1.DataVolumeSpec{
				Source: &v1beta1.DataVolumeSource{
					Blank: &v1beta1.DataVolumeBlankImage{},
				},
				Storage: &v1beta1.StorageSpec{
					AccessModes: []corev1.PersistentVolumeAccessMode{
						corev1.ReadWriteOnce,
					},
					Resources: corev1.ResourceRequirements{
						Requests: corev1.ResourceList{
							corev1.ResourceStorage: resource.MustParse("1Gi"),
						},
					},
					StorageClassName: utils.To[string]("openebs-localpv"),
				},
			},
		},
	}
}

func createDomain(sys, data string) virtcorev1.DomainSpec {
	return virtcorev1.DomainSpec{
		CPU: &virtcorev1.CPU{
			Cores:   1,
			Sockets: 1,
			Threads: 1,
		},

		Memory: &virtcorev1.Memory{
			Guest: resource.NewQuantity(512*1024*1024, resource.BinarySI), // 512Mi = 512 * 1024 * 1024 bytes
		},
		Machine: &virtcorev1.Machine{
			Type: "q35",
		},
		Devices: virtcorev1.Devices{
			Disks: []virtcorev1.Disk{
				{
					Name: sys,
					DiskDevice: virtcorev1.DiskDevice{
						Disk: &virtcorev1.DiskTarget{
							Bus: virtcorev1.DiskBusVirtio,
						},
					},
				},
				{
					Name: data,
					DiskDevice: virtcorev1.DiskDevice{
						Disk: &virtcorev1.DiskTarget{
							Bus: virtcorev1.DiskBusSATA,
						},
					},
				},
			},
			Interfaces: []virtcorev1.Interface{
				{
					Name: "default",
					InterfaceBindingMethod: virtcorev1.InterfaceBindingMethod{
						Masquerade: &virtcorev1.InterfaceMasquerade{},
					},
				},
			},
		},
	}
}

func createVolume(sys, sysDV, data, dataDV, hostname, sshkey, username, pwd string) []virtcorev1.Volume {
	// Cloud-Init user data for SSH and other initial setup
	cloudInitUserData := `#cloud-config
hostname: %s
ssh_pwauth: True
disable_root: false
ssh_authorized_keys:
  - %s
chpasswd:
  list: |
    %s:%s
  expire: False
`

	return []virtcorev1.Volume{
		{
			Name: sys,
			VolumeSource: virtcorev1.VolumeSource{
				DataVolume: &virtcorev1.DataVolumeSource{
					Name: sysDV,
				},
			},
		},
		{
			Name: "",
			VolumeSource: virtcorev1.VolumeSource{
				HostDisk: &virtcorev1.HostDisk{
					Path: "",
					Type: "",
					Capacity: resource.Quantity{
						Format: "",
					},
					Shared: nil,
				},
				PersistentVolumeClaim: &virtcorev1.PersistentVolumeClaimVolumeSource{
					PersistentVolumeClaimVolumeSource: corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: "",
						ReadOnly:  false,
					},
					Hotpluggable: false,
				},
				CloudInitNoCloud: &virtcorev1.CloudInitNoCloudSource{
					UserDataSecretRef: &corev1.LocalObjectReference{
						Name: "",
					},
					UserDataBase64: "",
					UserData:       "",
					NetworkDataSecretRef: &corev1.LocalObjectReference{
						Name: "",
					},
					NetworkDataBase64: "",
					NetworkData:       "",
				},
				CloudInitConfigDrive: &virtcorev1.CloudInitConfigDriveSource{
					UserDataSecretRef: &corev1.LocalObjectReference{
						Name: "",
					},
					UserDataBase64: "",
					UserData:       "",
					NetworkDataSecretRef: &corev1.LocalObjectReference{
						Name: "",
					},
					NetworkDataBase64: "",
					NetworkData:       "",
				},
				Sysprep: &virtcorev1.SysprepSource{
					Secret: &corev1.LocalObjectReference{
						Name: "",
					},
					ConfigMap: &corev1.LocalObjectReference{
						Name: "",
					},
				},
				ContainerDisk: &virtcorev1.ContainerDiskSource{
					Image:           "",
					ImagePullSecret: "",
					Path:            "",
					ImagePullPolicy: "",
				},
				Ephemeral: &virtcorev1.EphemeralVolumeSource{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
						ClaimName: "",
						ReadOnly:  false,
					},
				},
				EmptyDisk: &virtcorev1.EmptyDiskSource{
					Capacity: resource.Quantity{
						Format: "",
					},
				},
				DataVolume: &virtcorev1.DataVolumeSource{
					Name:         "",
					Hotpluggable: false,
				},
				ConfigMap: &virtcorev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "",
					},
					Optional:    nil,
					VolumeLabel: "",
				},
				Secret: &virtcorev1.SecretVolumeSource{
					SecretName:  "",
					Optional:    nil,
					VolumeLabel: "",
				},
				DownwardAPI: &virtcorev1.DownwardAPIVolumeSource{
					Fields:      nil,
					VolumeLabel: "",
				},
				ServiceAccount: &virtcorev1.ServiceAccountVolumeSource{
					ServiceAccountName: "",
				},
				DownwardMetrics: &virtcorev1.DownwardMetricsVolumeSource{},
				MemoryDump: &virtcorev1.MemoryDumpVolumeSource{
					PersistentVolumeClaimVolumeSource: virtcorev1.PersistentVolumeClaimVolumeSource{
						PersistentVolumeClaimVolumeSource: corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: "",
							ReadOnly:  false,
						},
						Hotpluggable: false,
					},
				},
			},
		},
		{
			Name: data,
			VolumeSource: virtcorev1.VolumeSource{
				DataVolume: &virtcorev1.DataVolumeSource{
					Name: dataDV,
				},
			},
		},
		{
			Name: "cloudinitdisk",
			VolumeSource: virtcorev1.VolumeSource{
				CloudInitNoCloud: &virtcorev1.CloudInitNoCloudSource{
					UserData: fmt.Sprintf(cloudInitUserData, hostname, sshkey, username, pwd),
				},
			},
		},
	}
}

func createTemplates(ns, virtName string) *virtcorev1.VirtualMachineInstanceTemplateSpec {
	return &virtcorev1.VirtualMachineInstanceTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"kubevirt.io/size":   "small",
				"kubevirt.io/domain": virtName,
			},
		},
		Spec: virtcorev1.VirtualMachineInstanceSpec{
			Domain:   createDomain(genSysDVName(ns, virtName), genDataDVName(ns, virtName)),
			Networks: createNet(),
			Volumes: createVolume(genSysDVName(ns, virtName), genSysDVName(ns, virtName),
				genDataDVName(ns, virtName), genDataDVName(ns, virtName), "forrest",
				"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCeUouKrUixAa/CMTqep3Ekh6IRZwkz7cs2gs9N3WcQL1uwzN/EX8+0jfEKSxaTKv37q2fU3kN3NBlIcI1bYMpyLlH1XbojIQA3mPq4Xoey377WqmMdbQFplztF3iwWCnAIphFWTMvB8+5hl6tNbKMt4uTgzYMhAywS7nLJH14IrcCNB9qjpn0FQXKBefL7r/vGeawm9vl03+gMKWXS0Oxp8rogXixtrx10J06DDaq/xkjbpPdqZEDHLhMJpdgUpbJSSgbL59QV8y070J2XJpPKdfvRCOQPleaW7IxSbVTqO+Bd/LvVD9IFSbaJYPlsLQC4VcLwaEyE+tcOPt+ZDYqwOd/QdiBTICefkRWUYEoQQSYvZoST7SyX37EprFKeKSvd0nnYoUkoM2WGMMUl0fv2BwwmXBjnp1kaK1MCGIy2ozdsgOqFNR2l+z1YigP/ynk/pGFnxyZ1Es4QxIDyYqrHbrUiPTTaLQ44s4vmXoFin+ZBzIFVYaKiDcD1jbo8s1k= root@master01", "ubuntu", "ubuntu"),
		},
	}
}

func createNet() []virtcorev1.Network {
	return []virtcorev1.Network{
		{
			Name: "default",
			NetworkSource: virtcorev1.NetworkSource{
				Pod: &virtcorev1.PodNetwork{}, // 指定使用 Pod 网络
			},
		},
	}
}

func CreateDefaultVM(namespace string, virtName string) *virtcorev1.VirtualMachine {
	return &virtcorev1.VirtualMachine{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kubevirt.io/v1",
			Kind:       "VirtualMachine",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      virtName,
		},
		Spec: virtcorev1.VirtualMachineSpec{
			Running:             utils.To[bool](true),
			Template:            createTemplates(namespace, virtName),
			DataVolumeTemplates: createDataVolume(genSysDVName(namespace, virtName), genDataDVName(namespace, virtName)),
		},
	}
}

// namespace-virtname--sys-DV
func genSysDVName(namespace, virtname string) string {
	return fmt.Sprintf("%s-%s-sys-dv", namespace, virtname)
}

// namespace-virtname-data-DV
func genDataDVName(namespace, virtname string) string {
	return fmt.Sprintf("%s-%s-data-dv", namespace, virtname)
}
