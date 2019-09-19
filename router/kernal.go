package router

import (
	"GoFormat/app/middleware"

	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// 組合log基本資訊
	r.Use(middleware.WriteLog)

	// api route
	LoadBackendRouter(r)

	// 載入graphql router設定
	LoadGraphqlRouter(r)
}
