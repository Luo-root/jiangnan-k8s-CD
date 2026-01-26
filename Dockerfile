# 第一阶段：构建阶段
FROM docker.donglizhiyuan.com/library/golang:1.25-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 复制 go 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译命令改为指向正确的主程序位置
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -a -o app ./cmd/

# 第二阶段：运行阶段
FROM docker.donglizhiyuan.com/library/alpine:3.20

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
