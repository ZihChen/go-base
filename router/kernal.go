package router

import (
	"goformat/app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" || os.Getenv("ENV") == "develop-ae" {
		r.Use(middleware.WriteLog)
	}

	// api route
	LoadBackendRouter(r)
}
