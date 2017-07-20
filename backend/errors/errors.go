package errors

import (
	"massliking/backend/logger"
)

var INSTABOT_REQUEST_ERROR = BaseError{1010, "Request to instagram failed", ""}
var INSTABOT_LOGIN_ERROR = BaseError{1020, "Instagram login failed", ""}
var INSTABOT_BODY_PARSE_ERROR = BaseError{1030, "Error parse instagram response", ""}
var INSTABOT_FOLLOWERS_ERROR = BaseError{1040, "Error fetch followers list", ""}
var INSTABOT_SUSPECTED_ERROR = BaseError{1050, "Instagram account is suspected", ""}

var MODEL_INSTAGRAM_NOT_FOUND = BaseError{2010, "Instagram not found", ""}
var MODEL_INSTAGRAM_COLLECTION_NOT_FOUND = BaseError{2020, "Instagram collection not found", ""}
var MODEL_INSTAGRAM_NOT_CREATED = BaseError{2030, "Instagram not created", ""}
var MODEL_INSTAGRAM_NOT_UPDATED = BaseError{2040, "Instagram not updated", ""}
var MODEL_INSTAGRAM_NOT_DELETED = BaseError{2050, "Instagram not deleted", ""}
var MODEL_INSTAGRAM_UNDEFINED_ACTION = BaseError{2060, "Undefined action type", ""}
var MODEL_INSTAGRAM_LOGIN_ERROR = BaseError{2070, "Instagram login failed", ""}
var MODEL_INSTAGRAM_INFO_ERROR = BaseError{2080, "Instagram info fetching failed", ""}
var MODEL_INSTAGRAM_INACTIVE_ERROR = BaseError{2090, "Instagram account is inactive", ""}

var MODEL_USER_NOT_FOUND = BaseError{3010, "User not found", ""}
var MODEL_USER_NOT_CREATED = BaseError{3030, "User not created", ""}

var MODEL_CHANNEL_UNDEFINED_ACTION = BaseError{4010, "Undefined action type", ""}
var MODEL_CHANNEL_ACTION_ERROR = BaseError{4020, "Error during channel action execution", ""}
var MODEL_CHANNEL_UNDEFINED_TARGET = BaseError{4030, "Undefined target type", ""}
var MODEL_CHANNEL_TARGET_ERROR = BaseError{4040, "Error during channel queue filling", ""}
var MODEL_CHANNEL_NOT_FOUND = BaseError{4050, "Channel not found", ""}
var MODEL_CHANNEL_COLLECTION_NOT_FOUND = BaseError{4060, "Channel collection not found", ""}
var MODEL_CHANNEL_NOT_CREATED = BaseError{4070, "Channel not created", ""}
var MODEL_CHANNEL_NOT_UPDATED = BaseError{4080, "Channel not updated", ""}
var MODEL_CHANNEL_NOT_DELETED = BaseError{4090, "Channel not deleted", ""}
var MODEL_CHANNEL_ACTION_EMPTY = BaseError{4100, "Empty action", ""}

var UNKNOWN_ERROR = BaseError{9999, "Unknown error", ""}

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}

func (e BaseError) Error() string {
	return e.Message
}

func On(err error, berr BaseError) error {
	if err != nil {
		logger.Error.Println(err)
		return berr
	} else {
		return nil
	}
}
