package routers

import (
	"go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/orders", controllers.GetOrders)

	return r
}
