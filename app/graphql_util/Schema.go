package graphql_util

import "github.com/graphql-go/graphql"

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
