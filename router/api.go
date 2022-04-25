package router

import (
	"goformat/app/global/helper"
	"goformat/app/handler/pprof"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {

	pprof.Register(r, "/internal/debug/pprof/") // 性能

	api := r.Group("/api")

	// K8S Health Check
	api.GET("/healthz", func(c *gin.Context) {
		data := map[string]string{
			"service": os.Getenv("PROJECT_NAME"),
			"time":    time.Now().Format("2006-01-02 15:04:05 -07:00"),
		}

		c.JSON(http.StatusOK, data)
	})

	// 載入測試用API
	if helper.IsDeveloperEnv(os.Getenv("ENV")) {
		// Swagger
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}
