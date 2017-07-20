package workers

import (
	"time"

	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

const LIMITS_TIMEOUT = 5 * 60
const ACTION_FAILURE_TIMEOUT = 30
const ACTION_EMPTY_TIMEOUT = 5 * 60

func ChannelsPoolWorker(client *instabot.Client, i *models.Instagram, action string, pool *models.WorkersPool) {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/channel", "ChannelsPoolWorker", action, i.IdString())

	logInfo("Start worker")

	for true {
		logInfo("Check commands")
		select {
		case command := <-pool.ChannelCommandChs[action]:
			if command == WORKER_COMMAND_STOP {
				logInfo("Command stop received")
				return
			}
		default:
		}

		logInfo("Validate account state")
		err = i.ValidateState(client)
		if err != nil {
			logError("Validate account state", err)
			return
		}

		logInfo("Looking for active channels with action: " + action)
		channels, err := i.FindActiveChannels(action)
		if err != nil {
			logError("Looking for active channels with action: "+action, err)
			return
		}

		if len(channels) == 0 {
			logInfo("Waiting for the next channels search")
			time.Sleep(time.Duration(ACTION_EMPTY_TIMEOUT) * time.Second)
			continue
		}

		logInfo("Looking for account followings for filtration")
		followings, _ := client.GetTotalUserFollowings(i.Info.PK)

		logInfo("Iterating over channels, filling queues, running actions")
		for _, c := range channels {
			_ = FillQueue(c, client, followings)

			if RunAction(i, c, client) {
				time.Sleep(time.Duration(3600/i.ActionSpeed(action)) * time.Second)
			} else {
				time.Sleep(time.Duration(ACTION_FAILURE_TIMEOUT) * time.Second)
			}
		}
	}
}
