package workers

import (
	"massliking/backend/instabot"
	"massliking/backend/models"
)

const WORKER_COMMAND_STOP = "stop"
const WRITE_OP_CAPACITY = 100

func BuildPool() *models.WorkersPool {
	pool := &models.WorkersPool{}

	pool.WriteOpCh = make(chan *models.InstagramWriteOp, WRITE_OP_CAPACITY)
	pool.InstagramCommandCh = make(chan string)
	pool.StatsCommandCh = make(chan string)
	pool.ChannelCommandChs = map[string](chan string){}
	for _, action := range models.ACTIONS {
		pool.ChannelCommandChs[action] = make(chan string)
	}

	return pool
}

func StartPool(i *models.Instagram, client *instabot.Client) {
	if models.INSTAGRAM_REGISTRY[i.Id] != nil {
		StopPool(i)
	}

	pool := BuildPool()

	go InstagramPoolWorker(client, i, pool)
	go StatsPoolWorker(client, i, pool)
	for _, action := range models.ACTIONS {
		go ChannelsPoolWorker(client, i, action, pool)
	}

	models.INSTAGRAM_REGISTRY[i.Id] = pool
}

func StopPool(i *models.Instagram) {
	if models.INSTAGRAM_REGISTRY[i.Id] == nil {
		return
	}

	pool := models.INSTAGRAM_REGISTRY[i.Id]

	pool.InstagramCommandCh <- WORKER_COMMAND_STOP
	pool.StatsCommandCh <- WORKER_COMMAND_STOP
	for _, action := range models.ACTIONS {
		pool.ChannelCommandChs[action] <- WORKER_COMMAND_STOP
	}

	delete(models.INSTAGRAM_REGISTRY, i.Id)
}
