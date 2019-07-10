package router

import (
	"GoFormat/app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	r.Use(middleware.WriteLog)

	// 載入graphql router設定
	LoadGraphqlRouter(r)

	// 載入測試用API
	if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" {
		LoadTestRouter(r)
		LoadGraphiqlToolRouter(r)
	}
}
