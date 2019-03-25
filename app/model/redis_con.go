package model

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"fmt"
	"time"

	Redis "github.com/gomodule/redigo/redis"
)

// RedisConnect 與redis連線
func RedisConnect() (redis Redis.Conn, apiErr errorcode.APIError) {
	// 使用redis封裝的Dial進行tcp連接
	host := global.Config.Redis.RedisHost
	port := global.Config.Redis.RedisPort
	pwd := global.Config.Redis.RedisPwd

	// 組合連接資訊
	var connectionString = fmt.Sprintf("%s:%s", host, port)
	redis, err := Redis.Dial(
		"tcp",
		connectionString,
		Redis.DialPassword(pwd),
		Redis.DialConnectTimeout(5*time.Second), // 建立連線 time out 時間 5 秒
		Redis.DialReadTimeout(5*time.Second),    // 讀取資料 time out 時間 5 秒
		Redis.DialWriteTimeout(5*time.Second),   // 寫入資料 time out 時間 5 秒
	)

	if err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_CONNECT_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("REDIS_CONNECT_ERROR")
		return
	}

	return
}

// RedisConnectTest 檢查 Redis 機器是否可以連線
func RedisConnectTest() {
	// 檢查 Master 連線
	redis, apiErr := RedisConnect()
	if apiErr.ErrorCode != 0 {
		panic("REDIS CONNECT ERROR")
	}
	defer redis.Close()
}
