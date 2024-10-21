- 下载win iso文件
```bash
WGET -L -o win10.iso https://go.microsoft.com/fwlink/p/?LinkID=2195443&clcid=0x804&culture=zh-cn&country=CN
```
- 获取cdi中的uploadproxy的service ip
```bash
kubectl -n cdi get svc -l cdi.kubevirt.io=cdi-uploadproxy
```
- 端口转发到本地
```bash
k port-forward -n cdi svc/cdi-uploadproxy 4443:443
```

- Upload the ISO
 ```
virtctl image-upload dv win10-iso --uploadproxy-url=https://127.0.0.1:4443 --size=10Gi --image-path=./win-server.iso --storage-class=openebs-localpv --volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```
