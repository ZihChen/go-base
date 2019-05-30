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

RUN echo "*/5 *	* * *	/usr/sbin/logrotate /etc/logrotate.conf" >> /etc/crontabs/root

CMD ["sh", "-c", "crond; govendor sync; ENV=${ENV:-develop} fresh runner.conf"]

##### 說明書 #####
# bash-4.4# crond --help
# BusyBox v1.28.4 (2018-05-30 10:45:57 UTC) multi-call binary.
# Usage: crond -fbS -l N -d N -L LOGFILE -c DIR
#        -f      Foreground
#        -b      Background (default)
#        -S      Log to syslog (default)
#        -l N    Set log level. Most verbose 0, default 8
#        -d N    Set log level, log to stderr
#        -L FILE Log to FILE
#        -c DIR  Cron dir. Default:/var/spool/cron/crontabs 
# This runs cron in the foreground with loglevel 2

# CMD [ "crond", "-l", "2", "-f" ]