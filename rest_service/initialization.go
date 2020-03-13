package rest_service

import (
	"github.com/gin-gonic/gin"
	"sync"
)

func InitializeRestService(wg sync.WaitGroup) {
	defer wg.Done()
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Start serving the application
	router.Run()
}

func initializeRoutes(router *gin.Engine) {

	fiboService := ProvideFibonacciService()
	router.POST("/", fiboService.GetFibonacciAnswer)

}

