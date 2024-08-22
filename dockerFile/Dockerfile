# 使用官方的 Go 镜像作为构建阶段的基础镜像
FROM golang:1.22.4-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录的所有文件复制到工作目录中
COPY . .

# 编译 Go 程序
RUN go build -o myapp .

# 创建一个运行阶段的更小的镜像
FROM alpine:3.18

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/myapp .

CMD ["-fileName=inDockerfile"]

# 指定默认的命令
ENTRYPOINT ["./myapp"]
