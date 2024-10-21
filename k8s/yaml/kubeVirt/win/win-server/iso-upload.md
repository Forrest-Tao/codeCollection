- 端口转发到本地
```bash
k port-forward -n cdi svc/cdi-uploadproxy 4453:443
```

- Upload the ISO
 ```
virtctl image-upload dv win-server-iso --uploadproxy-url=https://127.0.0.1:4443 --size=6Gi --image-path=./win-server.iso --storage-class=openebs-localpv --volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```
