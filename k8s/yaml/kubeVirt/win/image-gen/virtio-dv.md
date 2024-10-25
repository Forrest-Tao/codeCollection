```bash
kubectl virt image-upload  --kubeconfig ~/.bareconfig.yaml -n kubevirt \
dv win10-consumer-virtio-template --uploadproxy-url=https://127.0.0.1:4443 --size=4Gi \
--image-path=./virtio-win.iso --storage-class=openebs-localpv-kubevirt  \
--volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```

```bash
kubectl virt image-upload  --kubeconfig ~/.bareconfig.yaml -n kubevirt \
pvc win10-consumer-virtio --uploadproxy-url=https://127.0.0.1:4443 --size=4Gi \
--image-path=./virtio-win.iso --storage-class=openebs-localpv-kubevirt  \
--volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```