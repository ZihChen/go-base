package repository

import (
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"GoFormat/app/model"
	"fmt"
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

// Exists 檢查key是否存在
func (*Redis) Exists(key string) (ok bool, apiErr errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	chkExisits, _ := r.Do("EXISTS", key)
	ok, err := redis.Bool(chkExisits, nil)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_CHECK_EXIST_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("REDIS_CHECK_EXIST_ERROR")
		return
	}

	return
}

// Set 存入redis值
func (*Redis) Set(key string, value interface{}, time int) (apiErr errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	if _, err := r.Do("SET", key, value, "EX", time); err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_INSERT_ERROR: %v", err))
		apiErr = errorcode.GetAPIError("REDIS_INSERT_ERROR")
		return
	}

	return
}

// Get 取出redis值
func (*Redis) Get(key string) (value string, apiErr errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	value, _ = redis.String(r.Do("GET", key))

	return
}

// Delete 刪除redis值
func (*Redis) Delete(key string) (err errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	if _, err := r.Do("DEL", key); err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_DELETE_ERROR: %v", err))
		return errorcode.GetAPIError("REDIS_DELETE_ERROR")
	}

	return
}

// Append 在相同key新增多個值
func (*Redis) Append(key string, value interface{}) (n interface{}, apiErr errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return nil, apiErr
	}
	defer r.Close()

	n, err := r.Do("APPEND", key, value)
	if err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_APPEND_ERROR: %v", err))
		return nil, errorcode.GetAPIError("REDIS_APPEND_ERROR")
	}

	return
}

// HashSet Hash方式存入redis值
func (*Redis) HashSet(hkey string, key interface{}, value interface{}, time int) (err errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	// 存值
	if _, err := r.Do("hset", hkey, key, value); err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_INSERT_ERROR: %v", err))
		return errorcode.GetAPIError("REDIS_INSERT_ERROR")
	}

	// 設置過期時間
	if _, err := r.Do("EXPIRE", hkey, time); err != nil {
		go helper.WarnLog(fmt.Sprintf("REDIS_SET_EXPIRE_ERROR: %v", err))
		return errorcode.GetAPIError("REDIS_SET_EXPIRE_ERROR")
	}

	return
}

// HashGet Hash方式取出redis值
func (*Redis) HashGet(hkey string, field interface{}) (value string, apiErr errorcode.APIError) {
	r, apiErr := model.RedisConnect()
	if apiErr.ErrorCode != 0 {
		return
	}
	defer r.Close()

	// 取值
	value, _ = redis.String(r.Do("HGET", hkey, field))

	return
}
