package router

import (
	"goformat/app/global/helper"
	"goformat/app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	if helper.IsDeveloperEnv(os.Getenv("ENV")) {
		r.Use(middleware.WriteLog)
	}

	// api route
	LoadBackendRouter(r)
}
