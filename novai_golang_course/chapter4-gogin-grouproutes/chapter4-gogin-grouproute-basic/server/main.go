package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define routing groups
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")

	// Routes for v1 group
	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from API v1!",
		})
	})

	// Routes for v2 group
	v2.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from API v2!",
		})
	})

	// Start server
	router.Run(":8080")
}
