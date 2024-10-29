
尝试给vm扩容cpu
```bash
root@master01:/home/ubuntu/yst# kubectl patch vm cirros-vm -nkubevirt --type='json' -p='[{"op": "replace", "path": "/spec/template/spec/domain/cpu/sockets", "value": 3}]'
```

```bash
kc get vmi cirros-vm -oyaml
...
spec:
  architecture: amd64
  domain:
    cpu:
      cores: 1
      maxSockets: 4
      model: host-model
      sockets: 3
      threads: 1
...

...

status:
  activePods:
    1e2ed61a-a805-4183-86ab-f9900856fc22: slave04
    5a241186-97fc-4aa6-8448-aae3cab54bf4: slave04
    5b7bfdc6-280d-446e-9318-14e6fab29a76: slave05
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2024-10-29T02:49:35Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: null
    status: "True"
    type: LiveMigratable
  currentCPUTopology:
    cores: 1
    sockets: 3
    threads: 1
  guestOSInfo: {}
...

```

ref:
https://kubevirt.io/user-guide/compute/cpu_hotplug/
