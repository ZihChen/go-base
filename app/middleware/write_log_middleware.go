package middleware

import (
	"GoFormat/app/global/helper"

	"github.com/gin-gonic/gin"
)

// WriteLog 執行任何router前，都會紀錄一筆access.log
func WriteLog(c *gin.Context) {
	// 組合Log基本資料(IP,Path,Method,Status)
	helper.ComposeLog(c)

	// 	寫access Log
	helper.AccessLog()

	c.Next()
}
