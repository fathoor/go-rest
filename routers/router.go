package routers

import (
	"go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func ServerRouter() {
	r := gin.Default()

	r.GET("/orders", controllers.GetOrders)
}
