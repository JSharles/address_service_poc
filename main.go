package main

import (
	c "address/controllers"
	"address/routes"
	"address/utils"

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
	utils.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/address", c.CreateAddress)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(":3001")
}
