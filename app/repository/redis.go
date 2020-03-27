package repository

import (
	"goformat/app/global"
	"goformat/app/global/helper"
	"goformat/internal/cache"
	"goformat/library/errorcode"
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

// Redis 存取值
type Redis struct{}

var redisSingleton *Redis
var redisOnce sync.Once

// RedisIns 獲得單例對象
func RedisIns() *Redis {
	redisOnce.Do(func() {
		redisSingleton = &Redis{}
	})
	return redisSingleton
}

// RedisPing 檢查Redis是否啟動
func RedisPing() {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("🔔🔔🔔 REDIS CONNECT ERROR: %v 🔔🔔🔔", err.Error())
	}
}

// Exists 檢查key是否存在
func (*Redis) Exists(key string) (ok bool, apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	chkExisits, _ := conn.Do("EXISTS", key)
	ok, err := redis.Bool(chkExisits, nil)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_CHECK_EXIST_ERROR", err.Error())

		return
	}

	return
}

// Set 存入redis值
func (*Redis) Set(key string, value interface{}, expiretime int) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", key, value, "EX", expiretime); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_INSERT_ERROR", err.Error())

		return
	}
	return
}

// Get 取出redis值
func (*Redis) Get(key string) (value string, apiErr errorcode.Error) {

	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		helper.ErrorHandle(global.WarnLog, "REDIS_GET_VALUE_ERROR", err.Error(), key)
	}

	return
}

// Delete 刪除redis值
func (*Redis) Delete(key string) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_DELETE_ERROR", err.Error())

		return
	}

	return
}

// Append 在相同key新增多個值
func (*Redis) Append(key string, value interface{}) (n interface{}, apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	n, err := conn.Do("APPEND", key, value)
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_APPEND_ERROR", err.Error())

		return
	}

	return
}

// HashSet Hash方式存入redis值
func (*Redis) HashSet(hkey string, key interface{}, value interface{}, time int) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	// 存值
	if _, err := conn.Do("hset", hkey, key, value); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_INSERT_ERROR", err.Error())

		return
	}

	// 設置過期時間
	if _, err := conn.Do("EXPIRE", hkey, time); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_SET_EXPIRE_ERROR", err.Error())

		return
	}

	return
}

// HashMSet Hash 方式存入多筆 redis 值
func (*Redis) HashMSet(hkey string, value map[string]interface{}, time int) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	// 存入值
	if _, err := conn.Do("HMSET", redis.Args{}.Add(hkey).AddFlat(value)...); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_HASH_MSET_VALUE_ERROR", err.Error(), hkey, value)
		return
	}

	// 設置過期時間
	if _, err := conn.Do("EXPIRE", hkey, time); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_SET_EXPIRE_ERROR", err.Error())

		return
	}

	return
}

// HashGet Hash方式取出redis值
func (*Redis) HashGet(hkey string, field interface{}) (value string, apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	// 取值
	value, err := redis.String(conn.Do("HGET", hkey, field))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_HASH_GET_VALUE_ERROR", err.Error(), hkey, field)
	}

	return
}

// HashGetALL Hash方式取出所有redis值
func (*Redis) HashGetALL(hkey string, data interface{}) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	v, err := redis.Values(conn.Do("HGETALL", hkey))
	if err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_HASH_GET_ALL_VALUE_ERROR", err.Error(), hkey)
		return
	}

	if err := redis.ScanStruct(v, data); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_HASH_GET_SCAN_STRUCT_ERROR", err.Error(), hkey)
		return
	}

	return
}

// HashDelete 刪除redis 對應hash表
func (*Redis) HashDelete(hKey string, key interface{}) (apiErr errorcode.Error) {
	RedisPool := cache.RedisPoolConnect()
	conn := RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("HDEL", hKey, key); err != nil {
		apiErr = helper.ErrorHandle(global.WarnLog, "REDIS_DELETE_ERROR", err.Error())

		return
	}
	return

}
