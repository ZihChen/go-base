package business

import (
	"GoFormat/app/global"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/repository"
	"fmt"
	"sync"
	"time"
)

// RedisBus 管理者Business
type RedisBus struct {
}

var redisSingleton *RedisBus
var redisOnce sync.Once

// RedisIns 獲得單例對象
func RedisIns() *RedisBus {
	redisOnce.Do(func() {
		redisSingleton = &RedisBus{}
	})
	return redisSingleton
}

// SetRedisKey 存值進入redis
func (a *RedisBus) SetRedisKey() (err errorcode.Error) {
	redis := repository.RedisIns()
	key := fmt.Sprintf("GoFormat:TestRedis")
	err = redis.Set(key, time.Now(), global.RedisDBExpire)
	if err != nil {
		return
	}

	return
}

// GetRedisValue 取 redis 值
func (a *RedisBus) GetRedisValue() (value string, err errorcode.Error) {
	redis := repository.RedisIns()
	key := fmt.Sprintf("GoFormat:TestRedis")
	value, err = redis.Get(key)
	if err != nil {
		return
	}

	return
}
