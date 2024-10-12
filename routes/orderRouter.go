package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/orders", controller.GetOrders())
	upcomingRoutes.GET("/orders/:id", controller.GetOrder())
	upcomingRoutes.POST("/orders", controller.CreateOrder())
	upcomingRoutes.PATCH("/orders/:id", controller.UpdateOrder())
}
