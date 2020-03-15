package users

import (
	"net/http"

	"github.com/sebagalan/bookstore_users-api/types"
)

//GetUser ...
func GetUser(c *types.ContextRequest) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

//CreateUser ...
func CreateUser(c *types.ContextRequest) {
	c.String(http.StatusNotImplemented, "implemented me!")
}
