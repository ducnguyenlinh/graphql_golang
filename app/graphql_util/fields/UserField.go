package fields

import (
	"github.com/ducnguyenlinh/graphql_golang/app/graphql_util/types"
	"github.com/ducnguyenlinh/graphql_golang/app/model"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var users = []model.User{
	{
		UserId:		uuid.New().String(),
		UserName:	"user01",
		Email:		"user01@gmail.com",
	},
	{
		UserId:		uuid.New().String(),
		UserName:	"user02",
		Email:		"user02@gmail.com",
	},
	{
		UserId:		uuid.New().String(),
		UserName:	"user03",
		Email:		"user03@gmail.com",
	},
}

// Get user list
var UserListField = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "Get user list",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return users, nil
	},
}

// Create user
var CreateUserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"userName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["userName"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := model.NewUser(userName, email)
		if err != nil {
			panic(err)
		}

		users = append(users, newUser)
		return newUser, nil
	},
}
