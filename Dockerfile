# 使用官方的Golang Docker镜像作为基础
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将主机的~/downloads目录映射到容器的/downloads目录
VOLUME /downloads

# 将当前目录下的所有文件复制到容器的/app目录下
COPY . .

# 构建Go应用
RUN go build -o file-downloader

# 设置容器启动时执行的命令
CMD ["./file-downloader"]

