package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"

	"massliking/backend/auth"
	"massliking/backend/config"
	"massliking/backend/logger"
	"massliking/backend/models"
	"massliking/backend/workers"
)

func init() {
	rand.Seed(time.Now().Unix())

	// Initialize config
	config.Init()

	// Initialize logger
	logger.Init(config.GetString("log_file"))

	// Initialize db connection
	models.Init(config.GetString("db_string"))

	// Initialize jwt
	auth.Init(
		config.GetString("jwt.realm"),
		config.GetString("jwt.secret"),
		config.GetInt("jwt.ttl"),
	)

	// Launch all accounts workers
	workers.StartAllInstagrams()
}

func stop() {
	// Close db connection
	models.Stop()

	// Close logger
	logger.Stop()
}

func main() {
	// Build Gin instance and append middleware
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger)

	// Build routes
	InitRoutes(engine)

	// Launch server
	StartServer(engine)
}
