package mysql_utils

import (
	"github/lutungp/bookstore_users-api/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
