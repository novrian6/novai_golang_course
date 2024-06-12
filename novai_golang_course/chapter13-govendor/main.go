package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	r := gin.Default()

	// Monitoring middleware
	r.Use(func(c *gin.Context) {
		start := time.Now()

		// Process the request
		c.Next()

		// Calculate the duration
		duration := time.Since(start)

		// Log the request duration
		fmt.Printf("Request %s to %s took %v\n", c.Request.Method, c.Request.URL.Path, duration)
	})

	// Define a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	r.Run(":8080")
}

