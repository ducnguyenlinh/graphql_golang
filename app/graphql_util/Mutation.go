package graphql_util

import (
	"github.com/ducnguyenlinh/graphql_golang/app/graphql_util/fields"
	"github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": fields.CreateUserField,
	},
})
