package routes

import (
	controller "restaurant-mgmt-go/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(upcomingRoutes *gin.Engine) {
	upcomingRoutes.GET("/invoices", controller.GetInvoices())
	upcomingRoutes.GET("/invoices/:id", controller.GetInvoice())
	upcomingRoutes.POST("/invoices", controller.CreateInvoice())
	upcomingRoutes.PATCH("/invoices/:id", controller.UpdateInvoice())
}
