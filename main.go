package main

import (
	"go-rest/config"
	"go-rest/routers"
)

func main() {
	// Database
	config.ConnectDB()

	// Server
	HOST := config.Config.SERVER_HOST
	PORT := config.Config.SERVER_PORT
	routers.StartServer().Run(HOST + ":" + PORT)
}
