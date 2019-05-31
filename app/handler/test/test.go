package test

import (
	"GoFormat/app/business"
	"GoFormat/app/global/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetRedisValue 測試 Redis 存值
// @Summary 測試 Redis 存值
// @description Redis Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/set_redis [POST]
func SetRedisValue(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	redisBus := business.RedisIns()
	if apiErr := redisBus.SetRedisKey(); apiErr != nil {
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		return
	}

	c.JSON(http.StatusOK, helper.Success(""))
}

// GetRedisValue 測試 Redis 取值
// @Summary 測試 Redis 取值
// @description Redis Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/get_redis [GET]
func GetRedisValue(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	redisBus := business.RedisIns()
	value, err := redisBus.GetRedisValue()
	if err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success(value))
}

// PingDBOnce Ping DB 測試
// @Summary Ping DB 測試
// @description DB Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/ping_db_once [GET]
func PingDBOnce(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	dbBus := business.DBIns()
	if err := dbBus.PingDBOnce(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("123"))
}

// PingDBSecond Ping DB 測試
// @Summary Ping DB 測試
// @description DB Pool 連線測試
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/ping_db_second [GET]
func PingDBSecond(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	dbBus := business.DBIns()
	if err := dbBus.PingDBSecond(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("456"))
}

// ErrorTest 測試錯誤發生時是否可以回傳正確的 logID
// @Summary 測試錯誤發生時是否可以回傳正確的 logID
// @description DB Pool 測試錯誤發生時是否可以回傳正確的 logID
// @Tags Test
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/error_task [GET]
func ErrorTest(c *gin.Context) {
	// 接不可預期的錯誤
	defer helper.CatchError(c)

	errBus := business.ErrIns()
	if err := errBus.GetErrorLog(); err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success("999"))
}
