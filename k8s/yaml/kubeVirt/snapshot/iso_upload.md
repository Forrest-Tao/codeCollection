- 端口转发
```bash
k port-forward -n cdi svc/cdi-uploadproxy 4443:443 
```


- 上传iso并创建pvc存储更目录数据
```bash
virtctl image-upload dv ubuntu-iso --uploadproxy-url=https://127.0.0.1:4443 --size=5Gi --image-path=./noble-server-cloudimg-amd64.img --storage-class=openebs-localpv --volume-mode=filesystem --access-mode=ReadWriteOnce --force-bind --insecure
```

- 检查数据盘pvc是否挂载成功
```bash
ubuntu@ubuntu-vm:~$ lsblk
NAME    MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
sda       8:0    0    6G  0 volume 
├─sda1    8:1    0    5G  0 part /
├─sda14   8:14   0    4M  0 part 
├─sda15   8:15   0  106M  0 part /boot/efi
└─sda16 259:0    0  913M  0 part /boot
sdb       8:16   0    1M  0 volume 
sdc       8:32   0   10G  0 volume   🏅
```
k 

```bash
#关闭vm
virtctl stop ubuntu-vm

#创建vm的snapshot
k apply -f snapshot.yaml

#开启vm
virtctl start ubuntu-vm

#对根目录做一些修改
#退出vm
#关闭vm

#利用restore，回滚之前的版本
#开启vm
#检查是否restore成功
```