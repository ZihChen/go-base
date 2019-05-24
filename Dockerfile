# 第一層基底
FROM golang:1.11.2-alpine

# 取參數
ARG ENV
ARG PROJECT_NAME

# go get 會用到
RUN apk add git \
    && apk add logrotate

COPY ./logrotate /etc/logrotate.d/$PROJECT_NAME

# 複製專案
COPY . /go/src/$PROJECT_NAME

# 指定專案工作路徑
WORKDIR /go/src/$PROJECT_NAME

# 安裝govendor + realize
RUN go get github.com/pilu/fresh \
    && go get github.com/kardianos/govendor

CMD ["sh", "-c", "govendor sync; ENV=${ENV:-develop} fresh runner.conf"]