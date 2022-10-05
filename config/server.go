package config

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	PORT := Config.SERVER_PORT
	r := gin.Default()

	r.Run(":" + PORT)
}
