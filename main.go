// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initRedisClient() // If you're using Redis
	r := gin.Default()

	// Set up routes for the Admin and Public APIs
	apiRoutes := r.Group("/api/v1")
	{
		apiRoutes.POST("/ad", createAd)
		apiRoutes.GET("/ad", filterAds)
	}

	// Start the server
	r.Run(":8080")
}
