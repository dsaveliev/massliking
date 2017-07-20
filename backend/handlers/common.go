package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	. "massliking/backend/errors"
	"massliking/backend/logger"
	"massliking/backend/models"
)

func OnError(err error, c *gin.Context) {
	logger.Error.Println(err)

	switch parsedError := err.(type) {
	case BaseError:
		c.JSON(http.StatusBadRequest, parsedError)
	default:
		c.JSON(http.StatusBadRequest, UNKNOWN_ERROR)
	}
}

func getUser(c *gin.Context) (*models.User, error) {
	user := &models.User{}

	username := c.Keys["userID"].(string)
	user, err := models.GetUser(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func getInstagram(c *gin.Context) (*models.User, *models.Instagram, error) {
	user := &models.User{}
	instagram := &models.Instagram{}

	user, err := getUser(c)
	if err != nil {
		return user, instagram, err
	}

	instagram_id, err := strconv.ParseInt(c.Param("instagram_id"), 10, 64)
	if err != nil {
		return user, instagram, err
	}

	instagram, err = user.GetInstagram(instagram_id)
	if err != nil {
		return user, instagram, err
	}

	return user, instagram, nil
}

func findInstagrams(c *gin.Context) (*models.User, []*models.Instagram, error) {
	user := &models.User{}
	instagrams := []*models.Instagram{}

	user, err := getUser(c)
	if err != nil {
		return user, instagrams, err
	}

	instagrams, err = user.FindInstagrams()
	if err != nil {
		return user, instagrams, err
	}

	return user, instagrams, nil
}

func getChannel(c *gin.Context) (*models.Instagram, *models.Channel, error) {
	instagram := &models.Instagram{}
	channel := &models.Channel{}

	_, instagram, err := getInstagram(c)
	if err != nil {
		return instagram, channel, err
	}

	id, err := strconv.ParseInt(c.Param("channel_id"), 10, 64)
	if err != nil {
		return instagram, channel, err
	}

	channel, err = instagram.GetChannel(id)
	if err != nil {
		return instagram, channel, err
	}

	return instagram, channel, nil
}

func findChannels(c *gin.Context) (*models.Instagram, []*models.Channel, error) {
	instagram := &models.Instagram{}
	channels := []*models.Channel{}

	_, instagram, err := getInstagram(c)
	if err != nil {
		return instagram, channels, err
	}

	channels, err = instagram.FindChannels()
	if err != nil {
		return instagram, channels, err
	}

	return instagram, channels, nil
}
