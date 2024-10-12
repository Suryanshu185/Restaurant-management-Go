package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/menus", controller.GetMenus())
	upcomingRoutes.GET("/menus/:id", controller.GetMenu())
	upcomingRoutes.POST("/menus", controller.CreateMenu())
	upcomingRoutes.PATCH("/menus/:id", controller.UpdateMenu())
}
