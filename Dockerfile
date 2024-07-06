FROM golang:1.20 AS builder

WORKDIR /app
COPY . /app

RUN git config -l

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o welive .

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 设置时区
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/welive .
COPY --from=builder /app/config .

ENTRYPOINT ["/app/dwelive"]
CMD ["-c", "/app/welive.yml"]