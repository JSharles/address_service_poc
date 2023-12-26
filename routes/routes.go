package routes

import (
	ctrl "address/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		routes := api.Group("/addresses")
		{
			routes.POST("/", ctrl.CreateAddress)
			routes.GET("/", ctrl.GetAddresses)
			routes.GET("/:id", ctrl.GetAddressByID)
			routes.PUT("/:id", ctrl.UpdateAddress)
			routes.DELETE("/:id", ctrl.DeleteAddress)
		}
	}
}
