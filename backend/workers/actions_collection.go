package workers

import (
	"math/rand"
	"time"

	. "massliking/backend/errors"
	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

func RunAction(i *models.Instagram, c *models.Channel, client *instabot.Client) error {
	var err error
	logInfo, logWarn, logError := logger.TaggedLoggers("workers/actions_collection", "RunAction", i.IdString(), c.IdString())

	logInfo("Checking channel queue size...")
	if len(c.Queue.Targets) == 0 {
		logWarn("Channel queue is empty")
		return MODEL_CHANNEL_ACTION_ERROR
	}

	logInfo("Checking channel schedule...")
	hour := time.Now().UTC().Hour()
	if hour < i.Hours.Min || hour > i.Hours.Max {
		logWarn("Channel out of schedule")
		return MODEL_CHANNEL_ACTION_ERROR
	}

	logInfo("Checking channel limits...")
	if !i.CheckLimits(c.Action) {
		logWarn("Channel limits exceeded")
		return MODEL_CHANNEL_ACTION_ERROR
	}

	logInfo("Running target action")
	pk := c.Queue.Targets[0]
	switch c.Action {
	case "like":
		err = RunActionLike(c, client, pk)
	case "comment":
		err = RunActionComment(c, client, pk)
	case "follow":
		err = RunActionFollow(c, client, pk)
	case "unfollow":
		err = RunActionUnfollow(c, client, pk)
	default:
		err = MODEL_CHANNEL_UNDEFINED_ACTION
	}

	if err != nil {
		logError("Running target action", err)
		logInfo("Skipping current target and move forward")
		err = c.Save(func(c *models.Channel) {
			c.Queue.Targets = c.Queue.Targets[1:]
		})
		if err != nil {
			logError("Skipping current target and move forward", err)
			return MODEL_CHANNEL_ACTION_ERROR
		}
	}

	logInfo("Updating channel limits...")
	err = i.UpdateLimits(c.Action)
	if err != nil {
		logError("Updating channel limits...", err)
		return MODEL_CHANNEL_ACTION_ERROR
	}
	logInfo("Updating channel queue...")
	err = c.Save(func(c *models.Channel) {
		c.Queue.Targets = c.Queue.Targets[1:]
		c.Queue.Leads = append(c.Queue.Leads, pk)
	})
	if err != nil {
		logError("Updating channel queue...", err)
		return MODEL_CHANNEL_ACTION_ERROR
	}

	logInfo(">>> Action completed successfully <<<")

	return nil
}

func RunActionLike(c *models.Channel, client *instabot.Client, pk int) error {
	var err error

	// Looking for posts of the last year
	t := time.Now().AddDate(-1, 0, 0).Unix()
	feed, err := client.GetUserFeed(pk, "", t)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if len(feed.Items) == 0 {
		return MODEL_CHANNEL_ACTION_EMPTY
	}

	mediaId := feed.Items[0].PK
	like, err := client.Like(mediaId)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if like.Status != "ok" {
		return MODEL_CHANNEL_ACTION_ERROR
	}

	return nil
}

func RunActionComment(c *models.Channel, client *instabot.Client, pk int) error {
	var err error

	if len(c.Comments) == 0 {
		return MODEL_CHANNEL_ACTION_EMPTY
	}

	// Looking for posts of the last year
	t := time.Now().AddDate(-1, 0, 0).Unix()
	feed, err := client.GetUserFeed(pk, "", t)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if len(feed.Items) == 0 {
		return MODEL_CHANNEL_ACTION_EMPTY
	}

	// Fetch the last post and random comment
	mediaId := feed.Items[0].PK
	text := c.Comments[rand.Intn(len(c.Comments))]

	comment, err := client.Comment(mediaId, text)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if comment.Status != "ok" {
		return MODEL_CHANNEL_ACTION_ERROR
	}

	return nil
}

func RunActionFollow(c *models.Channel, client *instabot.Client, pk int) error {
	var err error

	follow, err := client.Follow(pk)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if follow.Status != "ok" {
		return MODEL_CHANNEL_ACTION_ERROR
	}

	return nil
}

func RunActionUnfollow(c *models.Channel, client *instabot.Client, pk int) error {
	var err error

	unfollow, err := client.Unfollow(pk)
	if err != nil {
		return On(err, MODEL_CHANNEL_ACTION_ERROR)
	}

	if unfollow.Status != "ok" {
		return MODEL_CHANNEL_ACTION_ERROR
	}

	return nil
}
