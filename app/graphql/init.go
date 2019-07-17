package graphql

import (
	"GoFormat/app/graphql/schema/mutation"
	"GoFormat/app/graphql/schema/query"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func schemaRoot(c *gin.Context) (*graphql.Schema) {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery(),
			Mutation: rootMutation(c),
		},
	)
	return &schema

}

func rootQuery() (*graphql.Object) {
	//init root query
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root Query",
		//query here
		Fields: graphql.Fields{
			"hello": query.TestHello(),
		},
	})
	return rootQuery
}

func rootMutation(c *gin.Context) (*graphql.Object) {
	//init root mutation
	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Root Mutation",
		//mutation here
		Fields: graphql.Fields{
			"world": mutation.TestMutation(),
		},
	})
	return rootMutation
}

func GraphqlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := handler.New(&handler.Config{
			Schema:     schemaRoot(c),
			Pretty:     true,
			GraphiQL:   true,
			Playground: false,
		})
		h.ServeHTTP(c.Writer, c.Request)
	}
}
