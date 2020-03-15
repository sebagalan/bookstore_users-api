package users

import (
	"net/http"
	"strconv"

	"github.com/sebagalan/bookstore_users-api/domains/users"
	"github.com/sebagalan/bookstore_users-api/services"
	"github.com/sebagalan/bookstore_users-api/types"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

//GetUser ...
func GetUser(c *types.ContextRequest) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		restError := errors.NewBadRequestError("user id is not valid")
		c.JSON(restError.Status, restError)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

//CreateUser ...
func CreateUser(c *types.ContextRequest) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}
