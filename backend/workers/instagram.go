package workers

import (
	"time"

	"massliking/backend/instabot"
	"massliking/backend/logger"
	"massliking/backend/models"
)

const ACCOUNT_CHECK_TIME = 15 * 60

func InstagramPoolWorker(client *instabot.Client, i *models.Instagram, pool *models.WorkersPool) {
	var err error
	var writeOp *models.InstagramWriteOp
	logInfo, _, logError := logger.TaggedLoggers("workers/instagram", "InstagramPoolWorker", i.IdString())

	logInfo("Start worker")

	for {
		select {
		case writeOp = <-pool.WriteOpCh:
			err = writeOp.Instagram.SyncSave(writeOp.Callback)
			writeOp.ReadChannel <- &models.InstagramReadOp{Error: err}
		case command := <-pool.InstagramCommandCh:
			if command == WORKER_COMMAND_STOP {
				logInfo("Command stop received")
				return
			}
		case <-time.After(time.Second * ACCOUNT_CHECK_TIME):
			logInfo("Validate account state")
			err = i.ValidateState(client)
			if err != nil {
				logError("Validate account state", err)
				return
			}
		}
	}
}
