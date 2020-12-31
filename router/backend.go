package router

import (
	"goformat/app/handler/pprof"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {

	pprof.Register(r, "/api/debug/pprof/") // 性能

	api := r.Group("/api")

	// K8S Health Check
	api.GET("/healthz", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	// 載入測試用API
	if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" || os.Getenv("ENV") == "develop-ae" {
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}
