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

	user, getErr := services.UserService.GetUser(userID)
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

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

//UpdateUser ...
func UpdateUser(c *types.ContextRequest) {
	var user users.User

	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restError := errors.NewBadRequestError("user id is not valid")
		c.JSON(restError.Status, restError)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	isPartial := (c.Request.Method == http.MethodPatch)
	user.ID = userID
	result, updateErr := services.UserService.UpdateUser(isPartial, user)

	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

//DeleteUser ...
func DeleteUser(c *types.ContextRequest) {

	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restError := errors.NewBadRequestError("user id is not valid")
		c.JSON(restError.Status, restError)
		return
	}

	deleteErr := services.UserService.DeleteUser(userID)

	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

//Search ...
func Search(c *types.ContextRequest) {
	status := c.Query("status")

	users, err := services.UserService.FindByStatus(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users)

}

func Loggin(c *types.ContextRequest) {
	var request users.LogginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errBadRequest := errors.NewBadRequestError("bad request")
		c.JSON(errBadRequest.Status, errBadRequest)
		return

	}

	user, err := services.UserService.LogginRequest(request)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)

}
