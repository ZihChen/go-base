package router

import (
	"GoFormat/app/handler/admin"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// LoadTestRouter 僅限於開發站測試用路由控制
func LoadTestRouter(r *gin.Engine) {

	v1 := r.Group("/test")
	{
		// 刪除過期session
		v1.DELETE("/clear_session", admin.ClearExpiredSession)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
