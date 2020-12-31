# 第一層基底
# FROM golang:1.14.0-alpine
FROM nexus.cqgame.games/rd3/golang-tool:1.14.0-alpine

# 安裝 git
# go get fresh, grpc
RUN apk add git \
    && go get github.com/pilu/fresh \
    && go get google.golang.org/grpc

ARG ACCESS_TOKEN
ENV ACCESS_TOKEN=$ACCESS_TOKEN

RUN git config --global url."https://rd3-pkg:${ACCESS_TOKEN}@${GITLAB}".insteadOf "https://${GITLAB}"

# docker terminal 顯示 LOG
RUN mkdir -p /app/log/ \
    && ln -sf /dev/stdout /app/log/access.log \
    && ln -sf /dev/stdout /app/log/error.log
