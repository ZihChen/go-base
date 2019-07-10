package graphql

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)


var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

//init root query
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Root Query",
	//query here
	Fields: graphql.Fields{


	},
})

//init root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "Root Mutation",
	//mutation here
	Fields: graphql.Fields{

	},
})

func GraphqlHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}