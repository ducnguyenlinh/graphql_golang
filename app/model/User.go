package model

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
}

// constructor
func NewUser(name string, email string) (User, error) {
	// validate
	if name == "" {
		return User{"", "", ""}, errors.New("userName is empty")
	}
	if email == "" {
		return User{"", "", ""}, errors.New("email is empty")
	}

	return User{
		ID:		uuid.New().String(),
		Name:	name,
		Email:	email,
	}, nil
}
