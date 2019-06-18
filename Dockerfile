# 第一層基底
FROM golang:1.11.2-alpine

# 安裝 git, logrotate
# go get fresh
RUN apk add git logrotate supervisor \
    && go get github.com/pilu/fresh

RUN echo "*/5 *	* * *	/usr/sbin/logrotate /etc/logrotate.conf" >> /etc/crontabs/root