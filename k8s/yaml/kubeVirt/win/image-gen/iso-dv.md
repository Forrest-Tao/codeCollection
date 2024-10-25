```bash
k port-forward -n cdi svc/cdi-uploadproxy 4443:443
```


```bash
kubectl virt image-upload  --kubeconfig ~/.bareconfig.yaml -n kubevirt \
dv win10-consumer-iso-template --uploadproxy-url=https://127.0.0.1:4443 \
--size=50Gi --image-path=./WIN10-CONSUMER-22H2-DVD-CHINESE.iso --storage-class=openebs-localpv-kubevirt \
--volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```

```bash
kubectl virt image-upload  --kubeconfig ~/.bareconfig.yaml -n kubevirt \
dv win10-business-iso-template --uploadproxy-url=https://127.0.0.1:4443 \
--size=50Gi --image-path=./WIN10-BUSINESS-22H2-DVD-CHINESE.iso --storage-class=openebs-localpv-kubevirt \
--volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```


```bash
kubectl virt image-upload  --kubeconfig ~/.bareconfig.yaml -n kubevirt \
pvc win10-consumer-iso --uploadproxy-url=https://127.0.0.1:4443 \
--size=50Gi --image-path=./WIN10-CONSUMER-22H2-DVD-CHINESE.iso --storage-class=openebs-localpv-kubevirt \
--volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```