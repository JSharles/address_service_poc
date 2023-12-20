package main

import (
	"address/routes"
	"address/utils"

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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":3001")
}
