# 第一層基底
FROM nexus.cqgame.games/rd3/golang-tool:1.14.0-alpine AS build

# 取參數
ARG ACCESS_TOKEN
ARG PROJECT_NAME
ENV ACCESS_TOKEN=$ACCESS_TOKEN
ENV PROJECT_NAME=$PROJECT_NAME

## 設定私有庫網址
RUN go env -w GOPRIVATE=git.cchntek.com
RUN git config --global url."https://rd3-pkg:${ACCESS_TOKEN}@git.cchntek.com".insteadOf "https://git.cchntek.com"

# 載入翻譯包
RUN apk add git

# 複製原始碼
COPY . /go/src/${PROJECT_NAME}
WORKDIR /go/src/${PROJECT_NAME}

# 進行編譯(名稱為: runner)
RUN go build -o runner

#############################################################

# 最終運行golang 的基底
FROM nexus.cqgame.games/rd3/library-alpine:3.9.5

COPY --from=build /go/src/${PROJECT_NAME}/runner /app/runner
COPY ./env /app/env
WORKDIR /app

# 設定容器時區(美東)
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/America/Puerto_Rico /etc/localtime
    
RUN mkdir -p /app/log/
RUN ln -sf /dev/stdout /app/log/access.log \
    && ln -sf /dev/stdout /app/log/error.log

ENTRYPOINT [ "./runner" ]