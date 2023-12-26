package main

import (
	"address/database"
	h "address/handlers"
	"address/routes"

	_ "address/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Address Service API
// @description API for address management
// @version 1.0
// @host localhost:3001
// @BasePath /api
func main() {
	database.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/addresses", h.CreateAddress)
		v1.GET("/addresses", h.GetAddresses)
		v1.GET("/addresses/:id", h.GetAddressByID)
		v1.PUT("/addresses/:id", h.UpdateAddress)
		v1.DELETE("/addresses/:id", h.DeleteAddress)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(":3001")
}
