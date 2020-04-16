package model

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	UserId		string `json:"userId"`
	UserName	string `json:"userName"`
	Email       string `json:"email"`
}

// constructor
func NewUser(userName string, email string) (User, error) {
	// validate
	if userName == "" {
		return User{"", "", ""}, errors.New("userName is empty")
	}
	if email == "" {
		return User{"", "", ""}, errors.New("email is empty")
	}

	return User{
		UserId:      uuid.New().String(),
		UserName:    userName,
		Email:       email,
	}, nil
}
