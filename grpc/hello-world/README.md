
generate gPRC code
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

```bash
go mod tidy
```

```bash
go run server/main.go
```

```bash
go run client/main.go
```

