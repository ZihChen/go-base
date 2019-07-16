package router

import (
	"GoFormat/app/graphql"
	"github.com/gin-gonic/gin"
)

// LoadBackendRouter 路由控制
func LoadGraphqlRouter(r *gin.Engine) {
	//graphql route endpoint
	r.POST("/graphql", graphql.GraphqlHandler())
}

// LoadGraphiqlToolRouter GraphiQL工具
func LoadGraphiqlToolRouter(r *gin.Engine) {
	r.GET("/graphql", graphql.GraphqlHandler())
}

