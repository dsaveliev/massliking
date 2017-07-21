package workers

import (
	"time"

	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

const ACTION_EMPTY_TIMEOUT = 5 * 60
const ACTION_SUCCESS_TIMEOUT = 0
const ACTION_FAILURE_TIMEOUT = 5 * 60

func ChannelsPoolWorker(client *instabot.Client, i *models.Instagram, action string, pool *models.WorkersPool) {
	logInfo, _, _ := logger.TaggedLoggers("workers/channel", "ChannelsPoolWorker", action, i.IdString())

	logInfo("Start worker")
	timeout := ACTION_SUCCESS_TIMEOUT

	for true {
		select {
		case command := <-pool.ChannelCommandChs[action]:
			if command == WORKER_COMMAND_STOP {
				logInfo("Command stop received")
				return
			}
		case <-time.After(time.Second * time.Duration(timeout)):
			timeout = RunChannels(client, i, action)
		}
	}
}

func RunChannels(client *instabot.Client, i *models.Instagram, action string) int {
	var err error
	logInfo, _, logError := logger.TaggedLoggers("workers/channel", "RunChannels", action, i.IdString())

	logInfo("Validate account state")
	err = i.ValidateState(client)
	if err != nil {
		logError("Validate account state", err)
		return ACTION_FAILURE_TIMEOUT
	}

	logInfo("Looking for active channels with action: " + action)
	channels, err := i.FindActiveChannels(action)
	if err != nil {
		logError("Looking for active channels with action: "+action, err)
		return ACTION_FAILURE_TIMEOUT
	}

	if len(channels) == 0 {
		logInfo("Waiting for the next channels search")
		return ACTION_EMPTY_TIMEOUT
	}

	logInfo("Looking for account followings for filtration")
	followings, _ := client.GetTotalUserFollowings(i.Info.PK)

	logInfo("Iterating over channels, filling queues, running actions")
	for _, c := range channels {
		_ = FillQueue(c, client, followings)
		err = RunAction(i, c, client)
		if err != nil {
			return ACTION_FAILURE_TIMEOUT
		}
		time.Sleep(time.Duration(3600/i.ActionSpeed(action)) * time.Second)
	}

	return ACTION_SUCCESS_TIMEOUT
}
