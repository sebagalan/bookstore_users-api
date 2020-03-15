package users

import (
	"strings"

	"github.com/sebagalan/bookstore_users-api/types"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

//User ...
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Dto ..
type Dto interface {
	Validate() *types.RestError
}

//Validate ...
func (u *User) Validate() *types.RestError {

	email := strings.TrimSpace(strings.ToLower(u.Email))

	if email == "" {
		return errors.NewBadRequestError("email not valid")
	}

	return nil
}
