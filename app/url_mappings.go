package app

import (
	"github.com/sebagalan/bookstore_users-api/controllers/ping"
	"github.com/sebagalan/bookstore_users-api/controllers/users"
	"github.com/sebagalan/bookstore_users-api/types"
)

//MapUrls ...
func MapUrls(router *types.RouterEngine) {
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.POST("/users/loggin", users.Loggin)

}
