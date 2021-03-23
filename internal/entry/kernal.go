package entry

import (
	"fmt"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/internal/bootstrap"
	"goformat/internal/schedule"
	"goformat/internal/server"
	"os"
)

// Run 執行服務
func Run() {

	// 設定優雅結束程序[監聽]
	bootstrap.SetupGracefulSignal()

	// 取得欲開啟服務環境變數
	service := os.Getenv("SERVICE")

	// 啟動服務
	switch service {
	// 執行 http 服務
	case "http":
		server.Run()
	// 執行 cron 服務
	case "cron":
		schedule.Run()
	// 執行 grpc 服務
	case "grpc":
	// 本機環境執行兩種服務
	case "all":
		go schedule.Run()
		server.Run()
	default:
		_ = helper.ErrorHandle(global.FatalLog, fmt.Sprintf("[❌ Fatal❌ ] SERVICE IS NOT EXIST: %v", service), "")
		fmt.Println("[❌ Fatal❌ ] SERVICE IS NOT EXIST: ", service)
	}
}
