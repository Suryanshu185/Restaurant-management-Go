package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/orderItems", controller.GetOrderItems())
	upcomingRoutes.GET("/orderItems/:id", controller.GetOrderItem())
	upcomingRoutes.GET("/orderItems-order/:id", controller.GetOrderItemsByOrder())
	upcomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	upcomingRoutes.PATCH("/orderItems/:id", controller.UpdateOrderItem())
}
