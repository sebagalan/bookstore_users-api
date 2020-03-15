package app

import (
	"github.com/sebagalan/bookstore_users-api/controllers/ping"
	"github.com/sebagalan/bookstore_users-api/controllers/users"
	"github.com/sebagalan/bookstore_users-api/types"
)

//MapUrls ...
func MapUrls(router *types.RouterEngine) {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
