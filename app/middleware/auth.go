package middleware

import (
	"goformat/app/global"
	"goformat/app/global/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PProfAuth pprof 使用權限驗證
func PProfAuth(c *gin.Context) {
	permission := c.GetHeader("pprof")

	if permission == "" {
		apiErr := helper.ErrorHandle(global.WarnLog, "NOT_ALLOW_TO_USE_PPROF", "you must provide header", permission)
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	// 型態轉換
	allow, err := strconv.ParseBool(permission)
	if err != nil {
		apiErr := helper.ErrorHandle(global.WarnLog, "CHANGE_PARAMS_TYPE_FAIL", err.Error())
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	// 檢查是否符合權限
	if !allow {
		apiErr := helper.ErrorHandle(global.WarnLog, "NOT_ALLOW_TO_USE_PPROF", "")
		c.JSON(http.StatusOK, helper.Fail(apiErr))
		c.Abort()
		return
	}

	c.Next()
}
