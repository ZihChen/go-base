FROM evergreen.guardians.one/rd3/golang-tool:1.16.0-alpine AS build

# 載入翻譯包
RUN apk add git \
    && go get github.com/liuzl/gocc

# 複製原始碼
COPY . /go/src/goformat
WORKDIR /go/src/goformat

# 產生 swag api 文件
RUN swag init
# 進行編譯(名稱為：runner)
RUN go build -o runner


# 最終運行golang 的基底
FROM evergreen.guardians.one/rd3/library-alpine:3.9.5

COPY --from=build /go/src/goformat/runner /app/runner
COPY ./env /app/env

WORKDIR /app


# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime
    
RUN mkdir -p /app/log/ \
    && ln -sf /dev/stdout /app/log/access.log \
    && ln -sf /dev/stdout /app/log/error.log

ENTRYPOINT [ "./runner" ]