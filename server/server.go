package server

import (
	"latest/config"
	"latest/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func StartServer() Server {
	return Server{
		port:   "10000",
		server: gin.Default(),
	}
}

func (s Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Fatal(router.Run(":" + config.GetConfig().ServerPort))

}
