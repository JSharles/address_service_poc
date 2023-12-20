package routes

import (
	c "address/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		addressRoutes := api.Group("/addresses")
		{
			addressRoutes.POST("/", c.CreateAddress)
			addressRoutes.GET("/", c.GetAddresses)
			addressRoutes.GET("/:id", c.GetAddressByID)
			addressRoutes.PUT("/:id", c.UpdateAddress)
			addressRoutes.DELETE("/:id", c.DeleteAddress)
		}
	}
}
