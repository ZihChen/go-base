# 第一層基底
FROM nexus.cqgame.games/rd3/golang-tool:1.14.0-alpine AS build

# 複製原始碼
COPY . /go/src/goformat
WORKDIR /go/src/goformat

# 進行編譯(名稱為：goformat)
RUN go build -o goformat

# 最終運行golang 的基底
FROM nexus.cqgame.games/rd3/library-alpine:3.9.5

COPY --from=build /go/src/goformat/goformat /app/goformat
COPY ./env /app/env
WORKDIR /app

# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime
    
RUN mkdir -p /app/log/
RUN ln -sf /dev/stdout /app/log/goformat_access.log \
    && ln -sf /dev/stdout /app/log/goformat_error.log

ENTRYPOINT [ "./goformat" ]