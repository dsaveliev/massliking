package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"massliking/backend/models"
	"massliking/backend/workers"
)

func StartChannelHandler(c *gin.Context) {
	_, channel, err := getChannel(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StartChannel(channel)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channel)
}

func StopChannelHandler(c *gin.Context) {
	_, channel, err := getChannel(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = workers.StopChannel(channel)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channel)
}

func CreateChannelHandler(c *gin.Context) {
	_, instagram, err := getInstagram(c)
	if err != nil {
		OnError(err, c)
		return
	}

	creds := &models.ChannelCredentials{}
	c.BindJSON(creds)

	channel, err := instagram.CreateChannel(creds)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channel)
}

func FindChannelsHandler(c *gin.Context) {
	_, channels, err := findChannels(c)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channels)
}

func GetChannelHandler(c *gin.Context) {
	_, channel, err := getChannel(c)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channel)
}

func UpdateChannelHandler(c *gin.Context) {
	instagram, channel, err := getChannel(c)
	if err != nil {
		OnError(err, c)
		return
	}

	creds := &models.ChannelCredentials{}
	c.BindJSON(creds)

	channel, err = instagram.UpdateChannel(channel, creds)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, channel)
}

func DeleteChannelHandler(c *gin.Context) {
	instagram, channel, err := getChannel(c)
	if err != nil {
		OnError(err, c)
		return
	}

	err = instagram.DeleteChannel(channel)
	if err != nil {
		OnError(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
