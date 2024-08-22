### 基本操作
```bash
#使用yaml创建configMap
k apply -f configmap.yaml
#使用文件夹创建 cm
k create cm cm-demo1 --from-file=testcm

 #使用文件创建cm
 k create cm cm-demo2 --from-file=testcm/redis.conf

#使用字面量创建configMap
 kubectl create configmap cm-demo3 --from-literal=db.host=localhost --from-literal=db.port=3306

```


### 作为环境变量使用

```bash
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
  - name: my-container
    image: my-image
    env:
    - name: MY_ENV_VAR
      valueFrom:
        configMapKeyRef:
          name: my-config
          key: key1
```
**MY_ENV_VAR** 环境变量将会被设置为 ConfigMap my-config 中 key1 对应的值 value1。