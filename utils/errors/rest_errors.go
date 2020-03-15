package errors

import (
	"net/http"

	"github.com/sebagalan/bookstore_users-api/types"
)

//NewBadRequestError ...
func NewBadRequestError(message string) *types.RestError {
	return &types.RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError ...
func NewNotFoundError(message string) *types.RestError {
	return &types.RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
