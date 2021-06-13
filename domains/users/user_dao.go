package users

import (
	"fmt"

	"github.com/sebagalan/bookstore_users-api/datasources/users_db"
	"github.com/sebagalan/bookstore_users-api/logger"
	"github.com/sebagalan/bookstore_users-api/utils/errors"
	"github.com/sebagalan/bookstore_users-api/utils/errors/date_utils"
	"github.com/sebagalan/bookstore_users-api/utils/errors/mysql_utils"
)

const (
	StatusActive                = "active"
	insertUserQuery             = "INSERT into users(first_name, last_name, email, date_created, status, password) values (?,?,?,?,?,?);"
	getOneUser                  = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=? LIMIT 1;"
	updateUserQuery             = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	deleteUserQuery             = "DELETE FROM users WHERE id=? LIMIT 1;"
	findByStatusQuery           = "SELECT id, first_name, last_name, email, date_created, status from users where status=?;"
	findByEmailAndPasswordQuery = "SELECT id, first_name, last_name, email, date_created, password, status from users where email = ?  LIMIT 1;"
)

//Dao ...
type Dao interface {
	Save() *errors.RestError
	Get() *errors.RestError
	Update() *errors.RestError
	Delete() *errors.RestError
	FindByStatus() *errors.RestError
	FindByEmailAndPassword() *errors.RestError
}

//Save ...
func (u *User) Save() *errors.RestError {
	stmt, err := users_db.UsersDb.Prepare(insertUserQuery)

	if err != nil {
		logger.Error("When prepare query", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	dateCreated := date_utils.GetNow()

	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, dateCreated, u.Status, u.Password)

	if saveErr != nil {
		logger.Error("When execute query", saveErr)
		return mysql_utils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("When return identifier for user", err)
		return mysql_utils.ParseError(err)
	}

	u.ID = userID
	u.DateCreated = dateCreated.Format(date_utils.APIDateFormat)
	return nil
}

//Get ...
func (u *User) Get() *errors.RestError {
	stmt, err := users_db.UsersDb.Prepare(getOneUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.ID)

	if getErr := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)

	}

	return nil
}

//Update ...
func (u *User) Update() *errors.RestError {
	stmt, err := users_db.UsersDb.Prepare(updateUserQuery)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.ID)

	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

//Delete ...
func (u *User) Delete() *errors.RestError {
	stmt, err := users_db.UsersDb.Prepare(deleteUserQuery)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(u.ID)

	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

//FindByStatus ...
func (u *User) FindByStatus() ([]User, *errors.RestError) {
	stmt, err := users_db.UsersDb.Prepare(findByStatusQuery)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, deleteErr := stmt.Query(u.Status)

	if deleteErr != nil {
		return nil, mysql_utils.ParseError(deleteErr)
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User

		if getErr := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
			return nil, mysql_utils.ParseError(getErr)

		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("not matching users with status %s", u.Status))
	}

	return results, nil
}

func (u *User) FindByEmailAndPassword() *errors.RestError {
	stmt, err := users_db.UsersDb.Prepare(findByEmailAndPasswordQuery)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Email)

	if getErr := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Password, &u.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}
