package main

import (
	"GoFormat/app/global"
	"GoFormat/app/global/helper"
	"GoFormat/app/model"
	"GoFormat/app/repository"
	_ "GoFormat/docs"
	"GoFormat/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var c *gin.Context

func main() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.WarnLog, fmt.Sprintf("[❌ Fatal❌ ]: %v", err), "")
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	// 開發時，console視窗不顯示有顏色的log
	gin.DisableConsoleColor()
	r := gin.Default()

	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	// 檢查 DB 機器服務
	model.DBPing()

	// 檢查 Redis 機器服務
	repository.RedisPing()

	// 背景
	// go task.Schedule()

	// 載入router設定
	router.RouteProvider(r)
	r.Run(":8080")
}
