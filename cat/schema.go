package cat

import (
	"errors"

	"github.com/graphql-go/graphql"
)

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
				Type: graphql.String,
			},
		},
	},
)

// queries
var catQueries = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CatQueries",
		Fields: graphql.Fields{
			"cats": &graphql.Field{
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						return GetCat(id), nil
					}
					return nil, nil
				},
			},
		},
	},
)

// mutations
var catMutations = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CatMutations",
		Fields: graphql.Fields{
			"addCat": &graphql.Field{
				Type: catSchema,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"color": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, ok := p.Args["name"].(string)
					if !ok {
						return nil, errors.New("name missing")
					}

					color, ok := p.Args["color"].(string)

					if !ok {
						return nil, errors.New("color missing")
					}

					return AddCat(name, color), nil
				},
			},
			"updateCat": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"color": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)

					if !ok {
						return nil, errors.New("id must be an integer")
					}
					name, ok := p.Args["name"].(string)
					if !ok {
						return nil, errors.New("name must be a string")
					}
					color, ok := p.Args["color"].(string)
					if !ok {
						return nil, errors.New("color must be a string")
					}

					return UpdateCat(id, name, color), nil
				},
			},
			"deleteCat": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)

					if !ok {
						return nil, errors.New("id missing")
					}

					return DeleteCat(id), nil
				},
			},
		},
	},
)

// used for Root Schema
var CatRootSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    catQueries,
		Mutation: catMutations,
	},
)
