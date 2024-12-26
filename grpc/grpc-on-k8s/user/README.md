

```bash
docker build -t grpc-on-k8s-user:v1 .
```

```bash
➜  user git:(main) ✗ kind load docker-image grpc-on-k8s-user:v1 --name kind
Image: "grpc-on-k8s-user:v1" with ID "sha256:e3d54502476c88ab69bb56b948bd27177539ff4891275aed070dcb59cc9f909a" not yet present on node "kind-control-plane", loading...
➜  user git:(main) ✗ ls                                                    
Dockerfile README.md  deploy     go.mod     go.sum     main.go    pb
➜  user git:(main) ✗ k apply -f deploy                                     
deployment.apps/user-service created
service/user-service created
```