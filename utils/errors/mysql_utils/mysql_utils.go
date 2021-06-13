package mysql_utils

import (
	"fmt"

	"github.com/go-sql-driver/mysql"

	"github.com/sebagalan/bookstore_users-api/utils/errors"
)

//ParseError ...
func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		return errors.NewInternalServerError(fmt.Sprintf("internal - %s", err.Error()))
	}

	return errors.NewBadRequestError(fmt.Sprintf("database error - %s", sqlErr.Error()))

}
