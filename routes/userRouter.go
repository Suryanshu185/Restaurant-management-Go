package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/users", controller.GetUsers())
	upcomingRoutes.GET("/users/:id", controller.GetUser())
	upcomingRoutes.POST("/users/signup", controller.SignUp())
	upcomingRoutes.POST("/users/login", controller.Login())
}
