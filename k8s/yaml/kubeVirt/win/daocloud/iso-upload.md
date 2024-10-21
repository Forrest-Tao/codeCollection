- 端口转发到本地
```bash
k port-forward -n cdi svc/cdi-uploadproxy 4453:443
```

- Upload the ISO
 ```
virtctl image-upload dv iso-win10-1 --uploadproxy-url=https://127.0.0.1:4453 --size=8Gi --image-path=./win-server.iso --storage-class=openebs-localpv --volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```
