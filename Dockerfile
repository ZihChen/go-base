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

# 美東時間
# RUN apk --no-cache add tzdata  && \
#     ln -sf /usr/share/zoneinfo/America/New_York /etc/localtime && \
#     echo "America/Los_Angeles" > /etc/timezone 

# 安裝govendor + fresh
RUN go get github.com/pilu/fresh \
    && go get github.com/kardianos/govendor

CMD ["sh", "-c", "govendor sync; ENV=${ENV:-develop} fresh runner.conf"]