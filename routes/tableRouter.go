package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/tables", controller.GetTables())
	upcomingRoutes.GET("/tables/:id", controller.GetTable())
	upcomingRoutes.POST("/tables", controller.CreateTable())
	upcomingRoutes.PATCH("/tables/:id", controller.UpdateTable())
}
