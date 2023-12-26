package main

import (
	c "address/controllers"
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

	v1 := r.Group("/api/v1")
	{
		// Swagger documentation route for CreateAddress
		v1.POST("/address", c.CreateAddress)

		// Your other routes go here...

		// Swagger UI route
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(":3001")
}
