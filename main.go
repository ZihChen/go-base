package main

import (
	"goformat/app/global"
	"goformat/app/repository"
	"goformat/internal/database"
	"goformat/internal/entry"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化動作
func init() {
	// 載入環境設定(所有動作須在該func後執行)
	global.Start()

	// 檢查 DB 機器服務
	database.DBPing()

	// 自動建置 DB + Table
	if os.Getenv("ENV") == "local" {
		database.CheckTableIsExist()
	}

	// 檢查 Redis 機器服務
	repository.RedisPing()

	// 設定程式碼 timezone
	os.Setenv("TZ", "America/Puerto_Rico")
}

func main() {
	entry.Run()
}
