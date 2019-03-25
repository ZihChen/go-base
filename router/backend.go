package router

import (
	"GoFormat/app/handler/admin"
	"GoFormat/app/middleware"

	"github.com/gin-gonic/gin"
)

// LoadBackendRouter 路由控制
func LoadBackendRouter(r *gin.Engine) {
	backend := r.Group("/backend")
	{
		backend.POST("login", admin.Login)

		// 此router以下才會使用Auth middleware
		backend.Use(middleware.Auth)
		{
			backend.GET("/user_info", admin.GetUserInfo)
			backend.GET("/menu", admin.CategoryMenu)
			backend.PUT("logout", admin.Logout)
			backend.PUT("/update_password", admin.UpdatePassword)
		}

		v1 := backend.Group("/admin", middleware.SuperAuth)
		{
			v1.GET("", admin.Index)
			v1.POST("/register", admin.Register)
			v1.PUT("/edit_admin/:id", admin.EditAdmin)
			v1.PUT("/reset_password/:id", admin.ResetPassword)
			v1.DELETE("/delete_admin/:id", admin.DeleteAdmin)
		}
	}
}
