package users

import (
	"fmt"
	"github/lutungp/bookstore_users-api/datasource/mysql/users_db"
	"github/lutungp/bookstore_users-api/utils/date_utils"
	"github/lutungp/bookstore_users-api/utils/errors"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInserUser   = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users where id=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d does not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInserUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, savErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if savErr != nil {
		sqlErr, ok := savErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("error when try to save user: %s", savErr.Error()))
		}

		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when try to save user: %s", savErr.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when try to get last insert id: %s", err.Error()))
	}
	user.Id = userId

	return nil
}
