package services

import (
	"github.com/sebagalan/bookstore_users-api/domains/users"
	"github.com/sebagalan/bookstore_users-api/types"
)

//CreateUser ...
func CreateUser(user users.User) (*users.User, *types.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//GetUser ...
func GetUser(userID int64) (*users.User, *types.RestError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}
