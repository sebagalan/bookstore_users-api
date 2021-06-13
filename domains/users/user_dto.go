package users

import (
	"strings"

	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

//User ...
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type RegisterUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//Dto ..
type Dto interface {
	Validate() *errors.RestError
}

//Validate ...
func (u *User) Validate() *errors.RestError {

	email := strings.TrimSpace(strings.ToLower(u.Email))

	if email == "" {
		return errors.NewBadRequestError("email not valid")
	}

	return nil
}
