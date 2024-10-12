package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/food", controller.GetFoods())
	upcomingRoutes.GET("/food/:id", controller.GetFood())
	upcomingRoutes.POST("/food", controller.CreateFood())
	upcomingRoutes.PATCH("/food/:id", controller.UpdateFood())
}
