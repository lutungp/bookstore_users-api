package services

import (
	"github/lutungp/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
