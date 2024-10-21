- 端口转发到本地
```bash
k port-forward -n cdi svc/cdi-uploadproxy 4443:443
```

- Upload the ISO
 ```
virtctl image-upload dv iso-win10 --uploadproxy-url=https://127.0.0.1:4443 --size=10Gi --image-path=./win10.iso --storage-class=openebs-localpv --volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```



