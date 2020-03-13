package rest_service

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"sync"
)

func InitializeRestService(wg sync.WaitGroup) {
	defer wg.Done()
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	router := setupRouter()
	// Start serving the application
	router.Run()
}

func setupRouter() *gin.Engine {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Initialize the routes
	close := initializeRoutes(router)
	runtime.SetFinalizer(router,
		func(f *gin.Engine) {
			close()
		})

	return router
}

func initializeRoutes(router *gin.Engine) func() error {

	fiboService, close := ProvideFibonacciService()
	router.POST("/", fiboService.GetFibonacciAnswer)

	return close
}
