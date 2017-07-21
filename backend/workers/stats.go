package workers

import (
	"time"

	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

const STATS_RELOAD_TIME = 15 * 60
const RELOGIN_TIME = 3 * 60 * 60

func StatsPoolWorker(client *instabot.Client, i *models.Instagram, pool *models.WorkersPool) {
	logInfo, _, _ := logger.TaggedLoggers("workers/stats", "StatsPoolWorker", i.IdString())

	logInfo("Start worker")

	for true {
		logInfo("Check commands")
		select {
		case command := <-pool.StatsCommandCh:
			if command == WORKER_COMMAND_STOP {
				logInfo("Command stop received")
				return
			}
		case <-time.After(time.Second * time.Duration(STATS_RELOAD_TIME)):
			UpdateStats(client, i, pool)
		}
	}
}

func UpdateStats(client *instabot.Client, i *models.Instagram, pool *models.WorkersPool) {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/stats", "UpdateStats", i.IdString())

	logInfo("Validate account state")
	err = i.ValidateState(client)
	if err != nil {
		logError("Validate account state", err)
		return
	}

	if time.Since(client.LoggedAt) >= time.Second*time.Duration(RELOGIN_TIME) {
		logInfo("Relogin account to update session")
		client, err = instabot.Login(i.Username, i.Password)
		if err != nil {
			logError("Relogin account to update session", err)
			return
		}
	}

	logInfo("Update account info")
	err = i.UpdateInfo(client)
	if err != nil {
		logError("Update account info", err)
	}

	logInfo("Fetch account channels")
	channels, err := i.FindChannels()
	if err != nil {
		logError("Fetch account channels", err)
	}

	if len(channels) > 0 {
		logInfo("Fetch account followers")
		followers, err := client.GetTotalUserFollowers(i.Info.PK)
		if err != nil {
			logError("Fetch account followers", err)
		}

		// Compare leads and followers
		for _, c := range channels {
			counter := 0
			for _, l := range c.Queue.Leads {
				if followers[l] == true {
					counter += 1
				}
			}

			// ...and save the result
			err = c.Save(func(c *models.Channel) {
				c.FollowersCount = counter
			})
			if err != nil {
				logError("Channel not updated", err)
			}
		}
	}
}
