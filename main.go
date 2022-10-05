package main

import (
	"go-rest/config"
	"go-rest/routers"
)

func main() {
	// Database
	config.ConnectDB()

	// Server
	PORT := ":" + config.Config.SERVER_PORT
	routers.StartServer().Run(PORT)
}
