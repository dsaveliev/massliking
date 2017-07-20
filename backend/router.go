package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"massliking/backend/auth"
	"massliking/backend/config"
	. "massliking/backend/handlers"
)

func InitRoutes(engine *gin.Engine) {
	engine.LoadHTMLGlob("./static/*.html")

	// NGINX serve static for production
	if config.IsDevelop() {
		engine.Static("/js", "./static/js")
		engine.Static("/statics", "./static/statics")
		engine.Static("/css", "./static/css")
		engine.Static("/fonts", "./static/fonts")
	}

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	v1 := engine.Group("/api/v1")
	{
		v1.POST("/login", auth.JWT.LoginHandler)
		v1.POST("/signup", SignupHandler)

		user := v1.Group("/user")
		user.Use(auth.JWT.MiddlewareFunc())
		{
			user.GET("", GetUserHandler)
			user.GET("/refresh_token", auth.JWT.RefreshHandler)
		}

		instagram := v1.Group("/instagrams")
		instagram.Use(auth.JWT.MiddlewareFunc())
		{
			instagram.POST("", CreateInstagramHandler)
			instagram.GET("", FindInstagramsHandler)
			instagram.GET("/:instagram_id", GetInstagramHandler)
			instagram.GET("/:instagram_id/stop", StopInstagramHandler)
			instagram.GET("/:instagram_id/start", StartInstagramHandler)
			instagram.PUT("/:instagram_id", UpdateInstagramHandler)
			instagram.DELETE("/:instagram_id", DeleteInstagramHandler)
		}

		channel := v1.Group("/instagram/:instagram_id/channels")
		channel.Use(auth.JWT.MiddlewareFunc())
		{
			channel.POST("", CreateChannelHandler)
			channel.GET("", FindChannelsHandler)
			channel.GET("/:channel_id", GetChannelHandler)
			channel.GET("/:channel_id/stop", StopChannelHandler)
			channel.GET("/:channel_id/start", StartChannelHandler)
			channel.PUT("/:channel_id", UpdateChannelHandler)
			channel.DELETE("/:channel_id", DeleteChannelHandler)
		}
	}

	engine.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

}
