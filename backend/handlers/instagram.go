package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"massliking/backend/models"
	"massliking/backend/workers"
)

func StopInstagramHandler(c *gin.Context) {
	_, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StopInstagram(instagram)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagram)
}

func StartInstagramHandler(c *gin.Context) {
	_, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StartInstagram(instagram)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagram)
}

func CreateInstagramHandler(c *gin.Context) {
	user, err := getUser(c)
	if err != nil {
		OnError(err, c)
		return
	}

	creds := &models.InstagramCredentials{}
	c.BindJSON(creds)

	instagram, err := user.CreateInstagram(creds)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StartInstagram(instagram)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagram)
}

func FindInstagramsHandler(c *gin.Context) {
	_, instagrams, err := findInstagrams(c)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagrams)
}

func GetInstagramHandler(c *gin.Context) {
	_, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagram)
}

func UpdateInstagramHandler(c *gin.Context) {
	user, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	creds := &models.InstagramCredentials{}
	c.BindJSON(creds)

	instagram, err = user.UpdateInstagram(instagram, creds)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, instagram)
}

func DeleteInstagramHandler(c *gin.Context) {
	user, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StopInstagram(instagram)
	if err != nil {
		OnError(err, c)
		return
	}

	err = user.DeleteInstagram(instagram)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
