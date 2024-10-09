- 利用deployment创建三个pod
- 利用service暴露他们
- 最终看是否为负载均衡式的服务


利用kind创建集群
```bash
kind create cluster --config ./kind-config.yaml
```

构建镜像
```bash
docker build -f Dockerfile -t echo-server:1.0 .
```

推送镜像到kind集群
```
kind load docker-image echo-server:1.0
```

创建deployment和service
```bash
k create -f ./yaml
```

结果
```bash
➜  service git:(main) ✗ k get pod -n echo-server -owide
NAME                                READY   STATUS    RESTARTS   AGE     IP           NODE           NOMINATED NODE   READINESS GATES
node-echo-server-64d96c4669-mtxcc   1/1     Running   0          3m20s   10.244.1.2   kind-worker3   <none>           <none>
node-echo-server-64d96c4669-pl7rx   1/1     Running   0          3m20s   10.244.2.2   kind-worker2   <none>           <none>
node-echo-server-64d96c4669-z6rl5   1/1     Running   0          3m20s   10.244.3.2   kind-worker    <none>           <none>
➜  service git:(main) ✗ curl http://localhost:30080/                                                                                                       
This pod is scheduled on node kind-worker
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker2
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker2
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker2
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker2
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker2
➜  service git:(main) ✗ curl http://localhost:30080/
This pod is scheduled on node kind-worker3
```
