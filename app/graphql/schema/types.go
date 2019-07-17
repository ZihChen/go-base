package schema

import "github.com/graphql-go/graphql"

//test type method
//type string,int,bool,float,ID
var HelloType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "World",
		Description: "Hello World api",
		Fields: graphql.Fields{
			"foo": &graphql.Field{
				Type: graphql.String,
			},
			"bar": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
