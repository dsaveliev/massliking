package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}
