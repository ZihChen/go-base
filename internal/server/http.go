package server

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/app/service"
	"goformat/app/task"
	"goformat/internal/bootstrap"
	"goformat/router"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

// RunHTTP HTTP 啟動 restful 服務
func RunHTTP() {
	defer func() {
		if err := recover(); err != nil {
			// 補上將err傳至telegram
			helper.ErrorHandle(global.FatalLog, fmt.Sprintf("[❌ Fatal❌ ] HTTP: %v", err), "")
			fmt.Println("[❌ Fatal❌ ] HTTP:", err)
		}
	}()

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

	// api gateway服務註冊
	out, _ := service.GateWayServiceRegister(global.Config.GrpcSetting.Name)
	fmt.Println(out)

	// 載入router設定
	router.RouteProvider(r)

	waitFinish := new(sync.WaitGroup)

	waitFinish.Add(1)
	go func(waitFinish *sync.WaitGroup) {
		defer waitFinish.Done()

		err := r.Run(":8080")
		if err != nil {
			helper.ErrorHandle(global.FatalLog, "SERVER_LISTEN_ERROR", err.Error())
			fmt.Println("[❌ Fatal❌ ] Server 建立監聽連線失敗:", err)
		}
	}(waitFinish)

	// 關閉優雅程序
	<-bootstrap.GracefulDown()

	waitFinish.Wait()
}
