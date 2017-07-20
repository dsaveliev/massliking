package workers

import (
	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

func StartAllInstagrams() {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StartAllInstagrams")

	logInfo("Find all accounts")
	instagrams, err := models.FindAllInstagrams()
	if err != nil {
		logError("Find all accounts", err)
	}

	for _, i := range instagrams {
		if i.State == models.INSTAGRAM_STATE_START {
			_ = StartInstagram(i)
		}
	}
}

func StartInstagram(i *models.Instagram) error {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StartInstagram", i.IdString())

	logInfo("Login account")
	client, err := instabot.Login(i.Username, i.Password)
	if err != nil {
		logError("Login account", err)
		return err
	}

	logInfo("Launch workers")
	StartPool(i, client)

	logInfo("Update account state")
	err = i.Save(func(i *models.Instagram) {
		i.State = models.INSTAGRAM_STATE_START
	})
	if err != nil {
		logError("Update account state", err)
		return err
	}

	logInfo("Update account info")
	err = i.UpdateInfo(client)
	if err != nil {
		logError("Update account info", err)
		return err
	}

	return nil
}

func StopInstagram(i *models.Instagram) error {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StopInstagram", i.IdString())

	logInfo("Stop all account channels")
	err = StopAllChannels(i)
	if err != nil {
		logError("Stop all account channels", err)
		return err
	}

	logInfo("Update account state")
	err = i.Save(func(i *models.Instagram) {
		i.State = models.INSTAGRAM_STATE_STOP
	})
	if err != nil {
		logError("Update account state", err)
		return err
	}

	logInfo("Stop workers")
	StopPool(i)

	return nil
}

func StopAllChannels(i *models.Instagram) error {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StopAllChannels", i.IdString())

	logInfo("Find all account channels")
	channels, err := i.FindChannels()
	if err != nil {
		logError("Find all account channels", err)
		return err
	}

	for _, channel := range channels {
		_ = StopChannel(channel)
	}

	return nil
}

func StartChannel(c *models.Channel) error {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StartChannel", c.IdString())

	logInfo("Update channel state")
	err = c.Save(func(c *models.Channel) {
		c.State = models.CHANNEL_STATE_START
	})
	if err != nil {
		logError("Update channel state", err)
		return err
	}

	return nil
}

func StopChannel(c *models.Channel) error {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/commands", "StopChannel", c.IdString())

	logInfo("Update channel state")
	err = c.Save(func(c *models.Channel) {
		c.State = models.CHANNEL_STATE_STOP
	})
	if err != nil {
		logError("Update channel state", err)
		return err
	}

	return nil
}
