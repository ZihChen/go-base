package router

import (
	"GoFormat/app/graphql"
	"os"

	"github.com/gin-gonic/gin"
)

// LoadGraphqlRouter 路由控制
func LoadGraphqlRouter(r *gin.Engine) {
	// 載入測試用API
	if os.Getenv("ENV") == "develop" || os.Getenv("ENV") == "local" {
		r.GET("/graphql", graphql.GraphqlHandler())
	}

	//graphql route endpoint
	r.POST("/graphql", graphql.GraphqlHandler())
}
