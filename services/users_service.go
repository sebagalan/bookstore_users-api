package services

import (
	"github.com/sebagalan/bookstore_users-api/domains/users"
	"github.com/sebagalan/bookstore_users-api/types"
)

//CreateUser ...
func CreateUser(user users.User) (*users.User, *types.RestError) {
	return &user, nil
}
