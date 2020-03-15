package users

import (
	"net/http"

	"github.com/sebagalan/bookstore_users-api/domains/users"
	"github.com/sebagalan/bookstore_users-api/services"
	"github.com/sebagalan/bookstore_users-api/types"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

//GetUser ...
func GetUser(c *types.ContextRequest) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

//CreateUser ...
func CreateUser(c *types.ContextRequest) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
	}
	c.JSON(http.StatusCreated, result)

}
