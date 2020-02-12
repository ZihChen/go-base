#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2019/01/14
# 說明： 啟動Golang的服務
#
# 備註：
#   

# 執行GoFormat的目錄,
WORK_PATH=$(dirname $(readlink -f $0))
# 執行各容器，須掛載的資料夾位置
VOLUME_PATH=$(dirname $(readlink -f $0))/../
# 專案名稱(取當前資料夾路徑最後一個資料夾名稱)
PROJECT_NAME=${WORK_PATH##*/}
# Log存放的目錄(預設local路徑)
LOG="/var/log/app/$PROJECT_NAME"
# 讀取圖片路徑(預設dev路徑)
IMG="$VOLUME_PATH/images"
# 環境變數
ENV="local"


# govendor path
GOVENDOR_PATH="$GOPATH/src/github.com/kardianos/govendor/"
# swagger path
SWAGGER_PATH="$GOPATH/src/github.com/swaggo/"


# 第一次clone專案須同步對外套件
if [ ! -d "$GOVENDOR_PATH" ]; then
    go get github.com/kardianos/govendor
fi

# 本機開發須安裝swagger + 初始化文件
if [ ! -d "$GOVENDOR_PATH" ]; then
    go get -u github.com/swaggo/swag/cmd/swag
fi

cd $WORK_PATH
govendor sync
swag init


#############################
#############################
docker network ls | grep "web_service" >/dev/null 2>&1
    if  [ $? -ne 0 ]; then
        docker network create web_service
    fi

echo "ENV=$ENV">.env
echo "LOG=$LOG">>.env
echo "IMG=$IMG">>.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

# 啟動容器服務
docker-compose up -d