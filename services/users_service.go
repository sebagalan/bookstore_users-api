package services

import (
	"github.com/sebagalan/bookstore_users-api/domains/users"
	"github.com/sebagalan/bookstore_users-api/utils/crypto_utils"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
	"github.com/sebagalan/bookstore_users-api/utils/errors/date_utils"
)

//UserService ...
var (
	UserService userServicesInterface = &usersServices{}
)

type usersServices struct{}

type userServicesInterface interface {
	CreateUser(users.User) (*users.RegisterUser, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.RegisterUser, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	FindByStatus(string) ([]users.User, *errors.RestError)
	LogginRequest(users.LogginRequest) (*users.RegisterUser, *errors.RestError)
}

//CreateUser ...
func (s *usersServices) CreateUser(user users.User) (*users.RegisterUser, *errors.RestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	hash, _ := crypto_utils.HashPassword(user.Password)

	println("CREATE USER", user.Password, hash)

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBSrting()
	user.Password = hash
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &users.RegisterUser{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Status:      user.Status,
		DateCreated: user.DateCreated,
	}, nil
}

//GetUser ...
func (s *usersServices) GetUser(userID int64) (*users.User, *errors.RestError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

//UpdateUser ...
func (s *usersServices) UpdateUser(isPartial bool, user users.User) (*users.RegisterUser, *errors.RestError) {

	current, err := UserService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName

		}
		if user.LastName != "" {
			current.LastName = user.LastName

		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return &users.RegisterUser{
		ID:        current.ID,
		FirstName: current.FirstName,
		LastName:  current.LastName,
		Email:     current.Email,
		Status:    current.Status,
	}, nil
}

//DeleteUser ...
func (s *usersServices) DeleteUser(userID int64) *errors.RestError {

	result := &users.User{ID: userID}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil

}

//FindByStatus ...
func (s *usersServices) FindByStatus(status string) ([]users.User, *errors.RestError) {

	result := &users.User{Status: status}
	return result.FindByStatus()

}

//LogginRequest ...
func (s *usersServices) LogginRequest(request users.LogginRequest) (*users.RegisterUser, *errors.RestError) {

	//intent, _ := crypto_utils.HashPassword(request.Password)
	result := &users.User{
		Email: request.Email,
	}

	if err := result.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	match := crypto_utils.CheckPasswordHash(request.Password, result.Password)

	println(request.Password, result.Password, match)

	if match {
		return &users.RegisterUser{
			ID:          result.ID,
			FirstName:   result.FirstName,
			LastName:    result.LastName,
			Email:       result.Email,
			Status:      result.Status,
			DateCreated: result.DateCreated,
		}, nil
	} else {
		err := errors.NewNotFoundError("user not found")
		return nil, err
	}
}
