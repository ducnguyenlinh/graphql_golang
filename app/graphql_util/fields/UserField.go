package fields

import (
	"errors"
	"github.com/ducnguyenlinh/graphql_golang/app/graphql_util/types"
	"github.com/ducnguyenlinh/graphql_golang/app/model"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var users = []model.User{
	{
		ID:		uuid.New().String(),
		Name:	"user01",
		Email:	"user01@gmail.com",
	},
	{
		ID:		uuid.New().String(),
		Name:	"user02",
		Email:	"user02@gmail.com",
	},
	{
		ID:		uuid.New().String(),
		Name:	"user03",
		Email:	"user03@gmail.com",
	},
}

// Get user list
var UserListField = &graphql.Field{
	Type:			graphql.NewList(types.UserType),
	Description:	"Get user list",
	Resolve:		func(p graphql.ResolveParams) (interface{}, error) {
		return users, nil
	},
}

// Find user by id
var UserField = &graphql.Field{
	Type:			types.UserType,
	Description:	"Get product by id",
	Args:			graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:	graphql.String,
		},
	},
	Resolve:		func(params graphql.ResolveParams) (interface{}, error) {
		params_id, ok := params.Args["id"]
		if ok {
			// Find user
			for _, user := range users {
				if user.ID == params_id {
					return user, nil
				}
			}
		}

		return nil, errors.New("user not found")
	},
}

// Create user
var CreateUserField = &graphql.Field{
	Type:			types.UserType,
	Description:	"Create new user",
	Args:			graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := model.NewUser(name, email)
		if err != nil {
			panic(err)
		}

		users = append(users, newUser)
		return newUser, nil
	},
}

// Update user by id
var UpdateUserField = &graphql.Field{
	Type:			types.UserType,
	Description:	"Update user by id",
	Args:			graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		params_id, _ := params.Args["id"].(string)
		params_name, nameOk := params.Args["name"].(string)
		params_email, emailOk := params.Args["email"].(string)

		user := model.User{}
		for i, u := range users {
			if u.ID == params_id {
				if nameOk {
					users[i].Name = params_name
				}
				if emailOk {
					users[i].Email = params_email
				}
				user = users[i]
				break
			}
		}
		return user, nil
	},
}

// Delete user by id
var DeleteUserField = &graphql.Field{
	Type:			types.UserType,
	Description:	"Delete user by id",
	Args:			graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve:		func(params graphql.ResolveParams) (interface{}, error) {
		params_id, _ := params.Args["id"].(string)
		user := model.User{}
		for i, u := range users {
			if u.ID == params_id {
				user = users[i]
				// Remove from user list
				users = append(users[:i], users[i+1:]...)
			}
		}

		return user, nil
	},
}
