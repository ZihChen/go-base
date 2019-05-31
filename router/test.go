package router

import (
	"GoFormat/app/handler/test"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadTestRouter 僅限於開發站測試用路由控制
func LoadTestRouter(r *gin.Engine) {

	v1 := r.Group("/test")
	{
		v1.GET("/get_redis", test.GetRedisValue)
		v1.POST("/set_redis", test.SetRedisValue)
		v1.GET("/ping_db_once", test.PingDBOnce)
		v1.GET("/ping_db_second", test.PingDBSecond)
		v1.GET("/error_task", test.ErrorTest)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
