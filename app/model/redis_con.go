package model

import (
	"GoFormat/app/global"
	"GoFormat/app/global/helper"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisPool 存放redis連線池的全域變數
var RedisPool *redis.Pool

func init() {
	fmt.Println("尚未連線前:", RedisPool)

	RedisPool = &redis.Pool{
		MaxIdle:     100,               // int 最大連線數
		MaxActive:   200,               // int 最大活耀連線數，默認為0不限制
		IdleTimeout: 300 * time.Second, // 連線過期時間，默認為0表示不做過期限制
		Wait:        true,              // 當連線超出限制數量後，是否等待到空閒連線釋放
		Dial: func() (redis.Conn, error) {
			// 使用redis封裝的Dial進行tcp連接
			host := global.Config.Redis.RedisHost
			port := global.Config.Redis.RedisPort
			pwd := global.Config.Redis.RedisPwd

			// 組合連接資訊
			var connectionString = fmt.Sprintf("%s:%s", host, port)
			redis, err := redis.Dial(
				"tcp",
				connectionString,
				redis.DialPassword(pwd),
				redis.DialConnectTimeout(5*time.Second), // 建立連線 time out 時間 5 秒
				redis.DialReadTimeout(5*time.Second),    // 讀取資料 time out 時間 5 秒
				redis.DialWriteTimeout(5*time.Second),   // 寫入資料 time out 時間 5 秒
			)

			if err != nil {
				go helper.WarnLog(fmt.Sprintf("REDIS_CONNECT_ERROR: %v", err))
				return nil, err
			}
			return redis, nil
		}, // 連接redis的函数
		TestOnBorrow: func(redis redis.Conn, t time.Time) error {
			// 每5秒ping一次redis
			if time.Since(t) < (5 * time.Second) {
				return nil
			}

			_, err := redis.Do("PING")
			if err != nil {
				go helper.WarnLog(fmt.Sprintf("REDIS_PING_ERROR: %v", err))
				return err
			}

			return nil
		}, // 定期對 redis server 做 ping/pong 測試

	}

	fmt.Println("連線後:", RedisPool)
}

// RedisPoolConnect 回傳連線池的 Redis 連線
func RedisPoolConnect() *redis.Pool {
	return RedisPool
}
