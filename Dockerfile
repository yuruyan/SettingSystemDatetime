# FROM dockerhub.world-machining.com/library/golang:1.22
FROM dockerhub.world-machining.com/library/golang:1.24.1-alpine3.21

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到Docker镜像中的工作目录
COPY . .

# 构建Go程序
RUN go build -ldflags="-w -s" -o myapp .

# 设置启动命令，运行构建好的Go程序
CMD ["./myapp"]
