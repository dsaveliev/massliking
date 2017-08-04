package instabot

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"

	. "massliking/backend/errors"
	"massliking/backend/logger"
)

type Client struct {
	PK           int
	Username     string
	Password     string
	Seed         string
	DeviceId     string
	GUUID        string
	Token        string
	RankToken    string
	Suspected    bool
	LastResponse []byte
	LoggedAt     time.Time
	Cookies      []*http.Cookie
}

func (c *Client) Init(username string, password string, cookies []*http.Cookie) {
	c.Cookies = cookies
	c.Username = username
	c.Password = password

	c.Seed = generateSeed(c.Username, c.Password)
	c.DeviceId = generateDeviceId(c.Seed)
	c.GUUID = generateUUID(true)

	c.Suspected = false

	c.LoggedAt = time.Now().AddDate(0, 0, -30)
}

func (c *Client) sendRequest(endpoint string, params string) (int, error) {
	var request *http.Request
	var err error
	logInfo, logWarn, logError := logger.TaggedLoggers("instabot/client", "sendRequest", c.Username, endpoint)

	if c.Suspected {
		logError("Client suspected, interrupting request", INSTABOT_SUSPECTED_ERROR)
		return -1, INSTABOT_SUSPECTED_ERROR
	}

	logInfo(">>> Request " + params)

	jar, _ := cookiejar.New(nil)
	jar.SetCookies(JAR_URL, c.Cookies)

	client := &http.Client{Jar: jar}

	if params != "" {
		request, err = http.NewRequest("POST", API_URL+endpoint, bytes.NewBuffer([]byte(params)))
		if err != nil {
			logError("Build POST request", err)
			return -1, err
		}
	} else {
		request, err = http.NewRequest("GET", API_URL+endpoint, nil)
		if err != nil {
			logError("Build GET request", err)
			return -1, err
		}
	}

	for k, v := range HEADERS {
		request.Header.Set(k, v)
	}

	response, err := client.Do(request)
	if err != nil {
		logError("Execute request", err)
		return -1, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logError("Read response body", err)
		return -1, err
	}

	response.Body.Close()

	logInfo("<<< Response: " + response.Status)
	logInfo("<<< Body: " + string(body))

	if response.StatusCode != http.StatusOK {
		logWarn("<<< Response: " + string(body))
		return response.StatusCode, INSTABOT_REQUEST_ERROR
	}

	if response.StatusCode == http.StatusTooManyRequests {
		c.Suspected = true
	}

	c.LastResponse = body
	c.Cookies = client.Jar.Cookies(JAR_URL)

	return response.StatusCode, nil
}

func (c *Client) findCookie(name string) string {
	for _, cookie := range c.Cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
