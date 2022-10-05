package main

import (
	"go-rest/config"
	"go-rest/routers"
)

func main() {
	// Database
	config.ConnectDB()

	// Server
	routers.ServerRouter()
	config.StartServer()
}
