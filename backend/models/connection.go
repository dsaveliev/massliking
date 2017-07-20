package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"

	"massliking/backend/logger"
)

var Engine *xorm.Engine

func Init(dbstring string) {
	var err error
	Engine, err = xorm.NewEngine("postgres", dbstring)
	if err != nil {
		logger.Error.Fatalf("Error opening db connection: %v", err)
	}

	Engine.SetLogger(xorm.NewSimpleLogger(logger.LogFile))

	err = Engine.Sync2(new(User))
	if err != nil {
		logger.Error.Fatalf("Error migrate model: User")
	}
	err = Engine.Sync2(new(Instagram))
	if err != nil {
		logger.Error.Fatalf("Error migrate model: Instagram")
	}
	err = Engine.Sync2(new(Channel))
	if err != nil {
		logger.Error.Fatalf("Error migrate model: Channel")
	}
}

func Stop() {
	Engine.Close()
}
