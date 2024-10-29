## 热插拔卷. OK

- 相关文章
  [Hotplug Volumes - KubeVirt user guide](https://kubevirt.io/user-guide/storage/hotplug_volumes/)

### 先决条件

开启HotplugVolumes.gate

```yaml
spec:
  certificateRotateStrategy: {}
  configuration:
    developerConfiguration:
      featureGates:
      - HotplugVolumes
```

### 创建一个dv

```bash
k apply -g example-volume-hotplug.yaml
```

**disk-type 的两种类型**

在 KubeVirt 中，disk-type 是用来指定磁盘设备的类型的参数，适用于热插拔（hotplugging）卷时的配置。它定义了热插拔卷的表现形式，以便在虚拟机（VM）内部的不同应用需求中使用不同的磁盘接

**disk**：

- 默认类型。如果未指定 disk-type，则默认使用 disk 类型。
- 这种类型将卷作为一个普通磁盘设备挂载到虚拟机中，适用于一般的存储需求，比如简单的数据存储和文件系统操作。
- 使用 disk 类型时，热插拔的卷将以虚拟磁盘的方式呈现在虚拟机中，并且支持标准的文件系统操作。

**lun (Logical Unit Number)**：

- 这种类型主要用来支持特殊的存储协议功能，比如 iSCSI 命令。
- 当 disk-type 设置为 lun 时，热插拔的卷允许更多低级操作，比如直接与存储设备进行数据块交互（即裸盘访问）。
- lun 类型通常用于数据库系统、需要直接访问块设备的应用或者涉及存储协议的场景，适合那些不需要文件系统管理、但需要与存储设备直接交互的应用。

### 增加零时卷

```yaml
root@master01:/home/ubuntu/yst# virtctl addvolume cirros-vm --volume-name=example-volume-hotplug-2 -nkubevirt
Successfully submitted add volume request to VM cirros-vm for volume example-volume-hotplug-2
root@master01:/home/ubuntu/yst#
```

in vm
```yaml
$ lsblk
NAME    MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
sda       8:0    0   1G  0 disk
vda     253:0    0  44M  0 disk
|-vda1  253:1    0  35M  0 part /
`-vda15 253:15   0   8M  0 part
$ [  759.227679] scsi 0:0:0:1: Direct-Access     QEMU     QEMU HARDDISK    2.5+ PQ: 0 ANSI: 5
[  759.232980] sd 0:0:0:1: Attached scsi generic sg1 type 0
[  759.237802] sd 0:0:0:1: Warning! Received an indication that the LUN assignments on this target have changed. The Linux SCSI layer does not automatical
[  759.247715] sd 0:0:0:1: [sdb] 4194304 512-byte logical blocks: (2.15 GB/2.00 GiB)
[  759.252843] sd 0:0:0:1: [sdb] Write Protect is off
[  759.254525] sd 0:0:0:1: [sdb] Write cache: enabled, read cache: enabled, doesn't support DPO or FUA
[  759.269244] sd 0:0:0:1: [sdb] Attached SCSI disk

$ lsblk
NAME    MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
sda       8:0    0   1G  0 disk
sdb       8:16   0   2G  0 disk
vda     253:0    0  44M  0 disk
|-vda1  253:1    0  35M  0 part /
`-vda15 253:15   0   8M  0 part
```

注意⚠️：发现 未设置 —persist参数时，即没有想要持久化，pvc也不会自动删除

### 删除永久卷

```yaml

root@master01:/home/ubuntu/yst# virtctl removevolume cirros-vm --volume-name=example-volume-hotplug -nkubevirt
Successfully submitted remove volume request to VM cirros-vm for volume example-volume-hotplug
```

in vm

```yaml
$ lsblk
NAME    MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
sda       8:0    0   1G  0 disk
vda     253:0    0  44M  0 disk
|-vda1  253:1    0  35M  0 part /
`-vda15 253:15   0   8M  0 part
$ [  447.237450] sd 0:0:0:0: [sda] Synchronizing SCSI cache

$ lsblk
NAME    MAJ:MIN RM SIZE RO TYPE MOUNTPOINT
vda     253:0    0  44M  0 disk
|-vda1  253:1    0  35M  0 part /
`-vda15 253:15   0   8M  0 part
```