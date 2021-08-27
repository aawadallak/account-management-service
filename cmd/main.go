package main

import (
	"latest/config"
	"latest/database"
	"latest/server"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Init()

	gin.SetMode(config.GetConfig().GinMode)

	webServer := server.StartServer()

	database.Init()

	webServer.Run()

}
