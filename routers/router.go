package routers

import (
	"go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/orders", controllers.GetOrders)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return r
}
