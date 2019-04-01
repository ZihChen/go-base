package main

import (
	"GoFormat/app/global"
	"GoFormat/app/global/helper"
	_ "GoFormat/docs"
	"GoFormat/router"
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var c *gin.Context

func main() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.FatalLog(err)
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	// 開發時，console視窗不顯示有顏色的log
	gin.DisableConsoleColor()
	r := gin.Default()

	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	// 檢查 DB 機器服務
	// model.DBConnectTest()

	// 檢查 Redis 機器服務
	// repository.RedisPing()

	// 背景
	// go task.Schedule()

	// 載入router設定
	router.RouteProvider(r)

	// Listen and Server in 0.0.0.0:8081
	server := endless.NewServer(":9999", r)
	server.BeforeBegin = func(add string) {
		log.Printf("Service Port %v, Actual pid is %d", add, syscall.Getpid())
		go helper.PidLog(syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
