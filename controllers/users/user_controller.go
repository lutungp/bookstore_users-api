package users

import (
	"fmt"
	"github/lutungp/bookstore_users-api/domain/users"
	"github/lutungp/bookstore_users-api/services"
	"net/http"

	"github/lutungp/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)
		fmt.Println(err.Error())

		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		fmt.Println(saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}
