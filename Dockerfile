# 使用官方的 Go 语言镜像作为基础镜像
FROM golang:1.18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 文件复制到工作目录
COPY go.mod ./

# 下载所有依赖并生成 go.sum 文件
RUN go mod tidy
RUN go mod download

# 将当前目录中的所有文件复制到工作目录
COPY . .

# 再次运行 go mod tidy 和 go mod download 以确保所有依赖项正确记录在 go.sum 文件中
RUN go mod tidy
RUN go mod download

# 构建 Go 应用程序
RUN go build -o main cmd/app/main.go

# 使用一个轻量级的基础镜像来运行应用程序
FROM alpine:latest

# 安装 ca-certificates 以便通过 HTTPS 进行连接
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建镜像复制二进制文件到当前镜像
COPY --from=builder /app/main .

# 运行应用程序
CMD ["./main"]