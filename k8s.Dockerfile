# 最終運行golang 的基底
FROM nexus.cqgame.games/rd3/library-alpine:3.9.5

COPY ./runner /app/runner
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