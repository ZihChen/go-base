package test

import (
	"GoFormat/app/business"
	"GoFormat/app/global/errorcode"
	"GoFormat/app/global/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetRedisValue 測試 Redis 存值
// @Summary 測試 Redis 存值
// @description Redis Pool 連線測試
// @Tags Admin
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/set_redis [POST]
func SetRedisValue(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

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
// @Tags Admin
// @Produce  json
// @Success 200 {object} structs.APIResult "成功"
// @Failure 400 {object} structs.APIResult "異常錯誤"
// @Router /test/get_redis [GET]
func GetRedisValue(c *gin.Context) {
	// 接Error
	defer func() {
		if err := recover(); err != nil {
			// 寫Fatal Log
			helper.FatalLog(err)

			// 回傳不可預期的錯誤
			apiErr := errorcode.GetAPIError(fmt.Sprintf("%v", err))
			c.JSON(http.StatusBadRequest, helper.Fail(apiErr))
		}
	}()

	redisBus := business.RedisIns()
	value, err := redisBus.GetRedisValue()
	if err != nil {
		c.JSON(http.StatusOK, helper.Fail(err))
		return
	}

	c.JSON(http.StatusOK, helper.Success(value))
}
