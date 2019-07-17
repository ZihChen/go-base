package schema

import "github.com/graphql-go/graphql"

//test type 方法
//type string,int,bool,float,ID
var HelloArgs = graphql.FieldConfigArgument{
	"foo": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	"bar": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}
