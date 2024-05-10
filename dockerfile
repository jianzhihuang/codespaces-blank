# 第一阶段：构建环境
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download   # 下载依赖
RUN go build -o main .  # 构建应用

# 第二阶段：运行环境
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/main .  # 从构建阶段复制 main 可执行文件
CMD ["./main"]  # 运行 main 程序
