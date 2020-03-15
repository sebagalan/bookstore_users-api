package users

import (
	"fmt"

	"github.com/sebagalan/bookstore_users-api/types"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
	"github.com/sebagalan/bookstore_users-api/utils/errors/date_utils"
)

var (
	usersDB = make(map[int64]*User)
)

//Dao ...
type Dao interface {
	Save() *types.RestError
	Get() *types.RestError
}

//Save ...
func (u *User) Save() *types.RestError {
	current := usersDB[u.ID]

	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", u.ID))
	}

	u.DateCreated = date_utils.GetNowSrting()

	usersDB[u.ID] = u
	return nil
}

//Get ...
func (u *User) Get() *types.RestError {
	result := usersDB[u.ID]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
	}

	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated

	return nil
}
