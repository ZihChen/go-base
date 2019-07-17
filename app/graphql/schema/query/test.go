package query

import (
	"GoFormat/app/global/structs"
	"GoFormat/app/graphql/schema"
	"github.com/graphql-go/graphql"
)

////test query 方法
func TestHello() (*graphql.Field) {
	var Testdata = graphql.Field{
		Name: "HelloWorld",
		Type: schema.HelloType,
		Args: schema.HelloArgs,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			foo, _ := p.Args["foo"].(string)
			bar, _ := p.Args["bar"].(string)
			var hello structs.HelloResponse
			hello.Foo = foo
			hello.Bar = bar
			return hello, nil
		},
	}
	return &Testdata

}
