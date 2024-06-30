FROM golang:1.20 AS builder

WORKDIR /app
COPY . /app

RUN export GOPRIVATE=e.coding.net \
    && go env -w GONOSUMDB=e.coding.net \
    && git config --global --add url."https://QVHXRIcDEG:40d0ee709d60f3bdce72d6a90a1058bc7e99d16d@e.coding.net".insteadOf "https://e.coding.net"

RUN git config -l

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o dh-scheduler .

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 设置时区
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/dh-scheduler .
COPY --from=builder /app/config .

ENTRYPOINT ["/app/dh-scheduler"]
CMD ["-c", "/app/settings.yml"]