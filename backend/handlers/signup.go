package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"massliking/backend/models"
)

func SignupHandler(c *gin.Context) {
	creds := &models.Credentials{}
	c.BindJSON(creds)

	user, err := models.CreateUser(creds)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}
