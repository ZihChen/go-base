package main

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/app/model"
	"goformat/app/repository"
	"goformat/app/task"
	"goformat/router"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var c *gin.Context

// 初始化動作
func init() {
	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	// 檢查 DB 機器服務
	model.DBPing()

	// 自動建置 DB + Table
	if os.Getenv("ENV") == "local" {
		model.CheckTableIsExist()
	}

	// 檢查 Redis 機器服務
	repository.RedisPing()

	// 設定程式碼 timezone
	os.Setenv("TZ", "America/Puerto_Rico")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.WarnLog, fmt.Sprintf("[❌ Fatal❌ ]: %v", err), "")
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	// 取得欲開啟服務環境變數
	service := os.Getenv("SERVICE")

	switch service {
	// 執行 http 服務
	case "http":
		runHTTP()
	// 本機環境執行兩種服務
	case "all":
		runHTTP()
	default:
		helper.ErrorHandle(global.FatalLog, fmt.Sprintf("[❌ Fatal❌ ] SERVICE IS NOT EXIST: %v", service), "")
		fmt.Println("[❌ Fatal❌ ] SERVICE IS NOT EXIST: ", service)
	}
}

// runHTTP HTTP 啟動服務
func runHTTP() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.FatalLog, fmt.Sprintf("[❌ Fatal❌ ] HTTP: %v", err), "")
			fmt.Println("[❌ Fatal❌ ] HTTP:", err)
		}
	}()

	// 開發時，console視窗不顯示有顏色的log
	gin.DisableConsoleColor()

	// 本機開發需要顯示 Gin Log
	var r *gin.Engine
	if os.Getenv("ENV") == "local" {
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	}

	// 背景
	go task.Schedule()

	// 載入router設定
	router.RouteProvider(r)
	r.Run(":8080")

	// 若有遇到需直接刪除 Pod 情形，為讓服務能 graceful-shutdown，可啟用以下代碼並將『r.Run(":8080")』註解即可
	//　https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	/*
		srv := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		go func() {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		// 設定 shutdown 的 timeout 時間 5s
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		// Shutdown 方法会等待活躍中連線的執行,但是等到 ctx 的 timeout 時間即刻 shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	*/
}
