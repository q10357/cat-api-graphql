package catservice

import "github.com/graphql-go/graphql"

var catSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Cat",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"color": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var catQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CatQueries",
		Fields: graphql.Fields{
			"cats": &graphql.Field {
				Type: graphql.NewList(catSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetCats(), nil
				},
			},
			"cat": &graphql.Field{
				Type: catSchema,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			}
		},
	}
)