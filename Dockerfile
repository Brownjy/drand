# 基于官方 Golang 镜像构建
FROM golang:1.20.1 AS builder

# 设置工作目录
WORKDIR /app

# 将 drand 源码复制到容器中
COPY . .

# 编译 drand 程序
RUN make install

# 生成 drand keypair
RUN drand generate-keypair --id default --scheme pedersen-bls-chained drand.kenlabs.org:443 > /root/.drand/group.toml

# 创建运行时镜像
FROM ubuntu:20.04

# 从构建阶段复制已编译的 drand 可执行文件
COPY --from=builder /go/bin/drand /usr/local/bin/

# 从构建阶段复制生成的 drand keypair 文件
COPY --from=builder /root/.drand/group.toml /root/.drand/

# 设置工作目录
WORKDIR /app

# 暴露端口
EXPOSE 8888
EXPOSE 8080
EXPOSE 9999

# 容器启动时执行 drand
CMD ["drand", "start", "--tls-disable", "--verbose", "--control", "8888", "--private-listen", "0.0.0.0:8080", "--metrics", "127.0.0.1:9999"]