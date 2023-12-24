package utils

import "github.com/gin-gonic/gin"

func SetupTestRouter() *gin.Engine {
	router := gin.Default()
	return router
}
