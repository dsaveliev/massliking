package workers

import (
	"math/rand"
	"strconv"
	"time"

	. "massliking/backend/errors"
	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

const MAX_TARGETS_PER_ONE_FILLING = 1000
const MIN_MEDIA_COUNT = 10

func FillQueue(c *models.Channel, client *instabot.Client, followings map[int]bool) error {
	var err error
	var users []*instabot.User
	var targets []int
	var leads map[int]bool
	logInfo, logWarn, logError := logger.TaggedLoggers("workers/targets_collection", "FillQueue", c.IdString())

	logInfo("Checking channel state...")
	if c.State != models.CHANNEL_STATE_START {
		logWarn("Exit due wrong channel state: " + c.State)
		return nil
	}

	logInfo("Checking channel queue size...")
	if len(c.Queue.Targets) > 0 {
		logWarn("Channel queue is not empty")
		return nil
	}

	logInfo("Looking for the next batch of targets")
	switch c.Target {
	case "followers":
		users, err = FillQueueByFollowers(c, client)
	case "subscriptions":
		users, err = FillQueueBySubscriptions(c, client)
	case "hashtag":
		users, err = FillQueueByHashtag(c, client)
	case "likes":
		users, err = FillQueueByLikes(c, client)
	case "comments":
		users, err = FillQueueByComments(c, client)
	default:
		err = MODEL_CHANNEL_UNDEFINED_TARGET
	}

	if err != nil {
		logError("Looking for the next batch of targets", err)
		_ = StopChannel(c)
		return err
	}

	logInfo("Targets amount: " + strconv.Itoa(len(users)))

	// Looking for already processed targets
	leads = c.Leads()

	//TODO: We should somehow limit batch size
	if len(users) > MAX_TARGETS_PER_ONE_FILLING {
		users = users[:MAX_TARGETS_PER_ONE_FILLING]
	}

	logInfo("Iterating over users, filtering and updating queue...")
	targets = filter(client, users, followings, leads)

	logInfo("Updating channel targets...")
	err = c.Save(func(c *models.Channel) {
		for t := range targets {
			c.Queue.Targets = append(c.Queue.Targets, t)
		}
	})
	if err != nil {
		logError("Updating channel targets...", err)
		_ = StopChannel(c)
		return err
	}

	return nil
}

func filter(client *instabot.Client, users []*instabot.User, followings map[int]bool, leads map[int]bool) []int {
	targets := []int{}

	for _, user := range users {
		// Accept only public accounts
		if user.IsPrivate == true {
			continue
		}

		// Decline our leads
		if leads[user.PK] == true {
			continue
		}

		// Decline our followers
		if followings[user.PK] == true {
			continue
		}

		// Some magick delay to avoid HTTP 429
		time.Sleep(time.Duration(1000+rand.Intn(1000)) * time.Millisecond)

		// Fetch account info
		response, err := client.SearchUsername(user.Username)
		if err != nil {
			continue
		}

		// Accept only real accounts, i.e. accounts with content
		info := response.UserInfo
		if info.MediaCount < MIN_MEDIA_COUNT {
			continue
		}

		// Append account id into leads collection
		leads[user.PK] = true

		targets = append(targets, user.PK)
	}

	return targets
}

func FillQueueByFollowers(c *models.Channel, client *instabot.Client) ([]*instabot.User, error) {
	var err error
	users := []*instabot.User{}

	info, err := client.SearchUsername(c.Value)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	pk := info.UserInfo.PK

	followers, err := client.GetUserFollowers(pk, c.Queue.MaxId)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	err = c.Save(func(c *models.Channel) {
		if followers.BigList != false {
			c.Queue.MaxId = followers.NextMaxID
		} else {
			c.Queue.MaxId = ""
			c.State = models.CHANNEL_STATE_EMPTY
		}
	})
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	return followers.Users, nil
}

func FillQueueBySubscriptions(c *models.Channel, client *instabot.Client) ([]*instabot.User, error) {
	var err error
	users := []*instabot.User{}

	info, err := client.SearchUsername(c.Value)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	pk := info.UserInfo.PK

	subscriptions, err := client.GetUserFollowings(pk, c.Queue.MaxId)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	err = c.Save(func(c *models.Channel) {
		if subscriptions.BigList != false {
			c.Queue.MaxId = subscriptions.NextMaxID
		} else {
			c.Queue.MaxId = ""
			c.State = models.CHANNEL_STATE_EMPTY
		}
	})
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	return subscriptions.Users, nil
}

func FillQueueByHashtag(c *models.Channel, client *instabot.Client) ([]*instabot.User, error) {
	var err error
	users := []*instabot.User{}

	feed, err := client.GetHashtagFeed(c.Value, c.Queue.MaxId)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	err = c.Save(func(c *models.Channel) {
		if feed.MoreAvailable != false {
			c.Queue.MaxId = feed.NextMaxID
		} else {
			c.Queue.MaxId = ""
			c.State = models.CHANNEL_STATE_EMPTY
		}
	})
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	for _, item := range feed.Items {
		users = append(users, item.User)
	}

	return users, nil
}

func FillQueueByLikes(c *models.Channel, client *instabot.Client) ([]*instabot.User, error) {
	users := []*instabot.User{}

	info, err := client.SearchUsername(c.Value)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	pk := info.UserInfo.PK

	// Looking for posts of the last year
	t := time.Now().AddDate(-1, 0, 0).Unix()
	feed, err := client.GetUserFeed(pk, c.Queue.MaxId, t)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	err = c.Save(func(c *models.Channel) {
		if feed.MoreAvailable != false {
			c.Queue.MaxId = feed.NextMaxID
		} else {
			c.Queue.MaxId = ""
			c.State = models.CHANNEL_STATE_EMPTY
		}
	})
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	//TODO: Skip likes pagination for now
	for _, item := range feed.Items {
		if item.LikeCount == 0 {
			continue
		}

		likes, err := client.GetMediaLikers(item.PK, "")
		if err != nil {
			continue
		}

		if likes.UserCount == 0 {
			continue
		}

		users = append(users, likes.Users...)
	}

	return users, nil
}

func FillQueueByComments(c *models.Channel, client *instabot.Client) ([]*instabot.User, error) {
	users := []*instabot.User{}

	info, err := client.SearchUsername(c.Value)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	pk := info.UserInfo.PK

	// Looking for posts of the last year
	t := time.Now().AddDate(-1, 0, 0).Unix()
	feed, err := client.GetUserFeed(pk, c.Queue.MaxId, t)
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	err = c.Save(func(c *models.Channel) {
		if feed.MoreAvailable != false {
			c.Queue.MaxId = feed.NextMaxID
		} else {
			c.Queue.MaxId = ""
			c.State = models.CHANNEL_STATE_EMPTY
		}
	})
	if err != nil {
		return users, On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	//TODO: Skip comments pagination for now
	for _, item := range feed.Items {
		if item.LikeCount == 0 {
			continue
		}

		comments, err := client.GetMediaComments(item.PK, "")
		if err != nil {
			continue
		}

		if comments.CommentCount == 0 {
			continue
		}

		for _, comment := range comments.Comments {
			users = append(users, comment.User)
		}
	}

	return users, nil
}
