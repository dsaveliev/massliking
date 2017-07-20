package instabot

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "massliking/backend/errors"
)

func Login(username string, password string) (*Client, error) {
	var err error

	client := Client{}
	client.Init(username, password, []*http.Cookie{})
	_, err = client.Login()

	if err != nil {
		return &client, On(err, MODEL_INSTAGRAM_LOGIN_ERROR)
	}

	return &client, nil
}

func (c *Client) SearchUsername(username string) (*SearchUsernameResponse, error) {
	var err error
	response := &SearchUsernameResponse{}

	_, err = c.sendRequest("users/"+username+"/usernameinfo/", "")
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) GetUserFollowers(pk int, maxid string) (*GetUserFollowersResponse, error) {
	var err error
	response := &GetUserFollowersResponse{}

	if maxid == "" {
		_, err = c.sendRequest("friendships/"+strconv.Itoa(pk)+"/followers/?rank_token="+c.RankToken, "")
	} else {
		_, err = c.sendRequest("friendships/"+strconv.Itoa(pk)+"/followers/?rank_token="+c.RankToken+"&max_id="+maxid, "")
	}
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) GetUserFollowings(pk int, maxid string) (*GetUserFollowingsResponse, error) {
	var err error
	response := &GetUserFollowingsResponse{}

	if maxid == "" {
		_, err = c.sendRequest("friendships/"+strconv.Itoa(pk)+"/following/?ig_sig_key_version="+SIG_KEY_VERSION+"&rank_token="+c.RankToken, "")
	} else {
		_, err = c.sendRequest("friendships/"+strconv.Itoa(pk)+"/following/?ig_sig_key_version="+SIG_KEY_VERSION+"&rank_token="+c.RankToken+"&max_id="+maxid, "")
	}
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) GetHashtagFeed(hashtag string, maxid string) (*GetHashtagFeedResponse, error) {
	var err error
	response := &GetHashtagFeedResponse{}

	if maxid == "" {
		_, err = c.sendRequest("feed/tag/"+hashtag+"/?rank_token="+c.RankToken+"&ranked_content=true", "")
	} else {
		_, err = c.sendRequest("feed/tag/"+hashtag+"/?max_id="+maxid+"&rank_token="+c.RankToken+"&ranked_content=true", "")
	}
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

// There is ids map for faster processing
func (c *Client) GetTotalUserFollowers(pk int) (map[int]bool, error) {
	var err error

	followers := map[int]bool{}
	nextMaxID := ""

	for true {
		response, err := c.GetUserFollowers(pk, nextMaxID)
		if err != nil {
			return followers, err
		}

		for _, u := range response.Users {
			followers[u.PK] = true
		}

		if response.BigList == false {
			return followers, nil
		}

		nextMaxID = response.NextMaxID
	}

	return followers, err
}

func (c *Client) GetTotalUserFollowings(pk int) (map[int]bool, error) {
	var err error

	followings := map[int]bool{}
	nextMaxID := ""

	for true {
		response, err := c.GetUserFollowings(pk, nextMaxID)
		if err != nil {
			return followings, err
		}

		for _, u := range response.Users {
			followings[u.PK] = true
		}

		if response.BigList == false {
			return followings, nil
		}

		nextMaxID = response.NextMaxID
	}

	return followings, err
}

func (c *Client) GetUserFeed(pk int, maxid string, minTimestamp int64) (*GetUserFeedResponse, error) {
	var err error
	response := &GetUserFeedResponse{}

	if maxid == "" {
		_, err = c.sendRequest("feed/user/"+strconv.Itoa(pk)+"/?min_timestamp="+strconv.FormatInt(minTimestamp, 10)+"&rank_token="+c.RankToken+"&ranked_content=true", "")
	} else {
		_, err = c.sendRequest("feed/user/"+strconv.Itoa(pk)+"/?max_id="+maxid+"&min_timestamp="+strconv.FormatInt(minTimestamp, 10)+"&rank_token="+c.RankToken+"&ranked_content=true", "")
	}
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) GetMediaLikers(mediaPk int, maxid string) (*GetMediaLikersResponse, error) {
	var err error
	response := &GetMediaLikersResponse{}

	if maxid == "" {
		_, err = c.sendRequest("media/"+strconv.Itoa(mediaPk)+"/likers/", "")
	} else {
		_, err = c.sendRequest("media/"+strconv.Itoa(mediaPk)+"/likers/?max_id="+maxid, "")
	}

	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) GetMediaComments(mediaPk int, maxid string) (*GetMediaCommentsResponse, error) {
	var err error
	response := &GetMediaCommentsResponse{}

	if maxid == "" {
		_, err = c.sendRequest("media/"+strconv.Itoa(mediaPk)+"/comments/", "")
	} else {
		_, err = c.sendRequest("media/"+strconv.Itoa(mediaPk)+"/comments/?max_id="+maxid, "")
	}
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) Follow(pk int) (*FollowResponse, error) {
	var err error
	response := &FollowResponse{}

	data := map[string]interface{}{
		"_uuid":      c.GUUID,
		"_uid":       c.PK,
		"user_id":    strconv.Itoa(pk),
		"_csrftoken": c.Token,
	}

	jsonData, _ := json.Marshal(data)
	_, err = c.sendRequest("friendships/create/"+strconv.Itoa(pk)+"/", generateSignature(jsonData))
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) Comment(mediaPk int, text string) (*CommentResponse, error) {
	var err error
	response := &CommentResponse{}

	data := map[string]interface{}{
		"_uuid":        c.GUUID,
		"_uid":         c.PK,
		"_csrftoken":   c.Token,
		"comment_text": text,
	}

	jsonData, _ := json.Marshal(data)
	_, err = c.sendRequest("media/"+strconv.Itoa(mediaPk)+"/comment/", generateSignature(jsonData))
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) Unfollow(pk int) (*UnfollowResponse, error) {
	var err error
	response := &UnfollowResponse{}

	data := map[string]interface{}{
		"_uuid":      c.GUUID,
		"_uid":       c.PK,
		"user_id":    strconv.Itoa(pk),
		"_csrftoken": c.Token,
	}

	jsonData, _ := json.Marshal(data)
	_, err = c.sendRequest("friendships/destroy/"+strconv.Itoa(pk)+"/", generateSignature(jsonData))
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) Like(mediaId int) (*LikeResponse, error) {
	var err error
	response := &LikeResponse{}

	data := map[string]interface{}{
		"_uuid":      c.GUUID,
		"_uid":       c.PK,
		"_csrftoken": c.Token,
		"media_id":   mediaId,
	}

	jsonData, _ := json.Marshal(data)
	_, err = c.sendRequest("media/"+strconv.Itoa(mediaId)+"/like/", generateSignature(jsonData))
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}
	return response, nil
}

func (c *Client) Login() (*LoginResponse, error) {
	var err error
	response := &LoginResponse{}

	_, err = c.sendRequest("si/fetch_headers/?challenge_type=signup&guid="+generateUUID(false), "")
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	data := map[string]interface{}{
		"phone_id":            generateUUID(true),
		"_csrftoken":          c.findCookie("csrftoken"),
		"username":            c.Username,
		"guid":                c.GUUID,
		"device_id":           c.DeviceId,
		"password":            c.Password,
		"login_attempt_count": "0",
	}

	jsonData, _ := json.Marshal(data)
	_, err = c.sendRequest("accounts/login/", generateSignature(jsonData))
	if err != nil {
		return response, INSTABOT_REQUEST_ERROR
	}

	err = json.Unmarshal(c.LastResponse, response)
	if err != nil {
		return response, INSTABOT_BODY_PARSE_ERROR
	}

	if response.Status != "ok" {
		return response, INSTABOT_LOGIN_ERROR
	}

	c.PK = response.LoggedInUser.PK
	c.RankToken = strconv.Itoa(c.PK) + "_" + c.GUUID
	c.Token = c.findCookie("csrftoken")
	c.LoggedAt = time.Now()

	return response, nil
}
