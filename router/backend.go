package router

import (
	"goformat/app/handler/pprof"
	"goformat/app/handler/test"
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
	{

		// K8S Health Check
		api.GET("/healthz", func(c *gin.Context) {
			c.AbortWithStatus(http.StatusOK)
		})

		// 載入測試用API
		if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" {
			v1 := api.Group("/test")
			{
				v1.GET("/get_redis", test.GetRedisValue)
				v1.POST("/set_redis", test.SetRedisValue)
				v1.GET("/ping_db_once", test.PingDBOnce)
				v1.GET("/ping_db_second", test.PingDBSecond)
				v1.GET("/error_task", test.ErrorTest)
			}

			// Swagger
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}
}
