package main

import (
	"os"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"massliking/backend/config"
	"massliking/backend/logger"
)

func SavePid(address string) {
	pid := syscall.Getpid()

	pidFile, err := os.OpenFile(config.GetString("pid_file"), os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		logger.Error.Fatalf("Error creating pid file: %v", err)
	}

	_, err = pidFile.Write([]byte(strconv.Itoa(pid)))
	if err != nil {
		logger.Error.Fatalf("Error write to pid file: %v", err)
	}

	pidFile.Close()

	logger.Info.Println("Listening and serving HTTP on", address)
	logger.Info.Printf("Actual pid is %d", pid)
}

func StartServer(engine *gin.Engine) {
	address := config.GetString("app_address")
	server := endless.NewServer(address, engine)

	server.BeforeBegin = SavePid

	err := server.ListenAndServe()
	if err != nil {
		logger.Error.Fatalf("Server interrupted: %v", err)
	}
}
