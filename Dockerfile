# 第一層基底
FROM golang:1.11.2-alpine

# 取參數
ARG ENV

# 複製專案
COPY . /go/src/GoFormat

# 指定專案工作路徑
WORKDIR /go/src/GoFormat

# RUN apk add --no-cache ca-certificates \
#         dpkg gcc git musl-dev

# go get 會用到
RUN apk add git

# 安裝govendor + realize
RUN go get github.com/tockins/realize \
    && go get github.com/kardianos/govendor

CMD ["sh", "-c", "govendor sync; ENV=${ENV:-develop} realize start"]