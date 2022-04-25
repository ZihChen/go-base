# 第一層基底
# FROM golang:1.16.0-alpine
FROM evergreen.guardians.one/rd3/golang-tool:1.16.0-alpine

# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime

# 安裝 git
# go get fresh, grpc
RUN apk add git \
    && go get github.com/pilu/fresh \
    && go get google.golang.org/grpc


# docker terminal 顯示 LOG
RUN mkdir -p /app/log/ \
    && ln -sf /dev/stdout /app/log/access.log \
    && ln -sf /dev/stdout /app/log/error.log
