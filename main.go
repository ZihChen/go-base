package main

import (
	"embed"
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/internal/cache"
	"goformat/internal/database"
	"goformat/internal/entry"
	"os"

	_ "goformat/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//go:embed env/*
var f embed.FS

// 初始化動作
func init() {
	// 載入環境設定(所有動作須在該func後執行)
	global.Start(f)

	// 檢查 DB 機器服務
	db := database.NewDbConnection()
	db.DBPing()

	// 自動建置 DB + Table
	if helper.IsLocalEnv(os.Getenv("ENV")) {
		db.CheckTableIsExist()
	}

	// 檢查 Redis 機器服務
	redis := cache.RedisIns()
	redis.RedisPing()

	// 設定程式碼 timezone
	os.Setenv("TZ", "America/Puerto_Rico")
}

func main() {
	entry.Run()
}
