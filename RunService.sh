#!/bin/bash
# 撰寫人員: Neil_Hsieh
# 撰寫日期：2019/01/14
# 說明： 啟動Golang的服務
#
# 備註：
#   

# 執行專案的目錄,
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
# swagger path
SWAGGER_PATH="$GOPATH/src/github.com/swaggo/"
# Gitlab Access Token(golang 使用私有庫套件，時需用到)
ACCESS_TOKEN="rmCquFPqfYsd9QrWTk_z"
# go module 存放路徑
GO_MOD_PATH="$GOPATH/pkg/mod"




# 本機開發須安裝swagger + 初始化文件
if [ ! -d "$GOVENDOR_PATH" ]; then
    echo "===== Swagger not exist, prepare to install ===="
    go get -u github.com/swaggo/swag/cmd/swag

fi

cd $WORK_PATH
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
echo "ACCESS_TOKEN"=$ACCESS_TOKEN>>.env
echo "GO_MOD_PATH"=$GO_MOD_PATH>>.env


# 啟動容器服務
docker-compose up -d