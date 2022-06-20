package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}
