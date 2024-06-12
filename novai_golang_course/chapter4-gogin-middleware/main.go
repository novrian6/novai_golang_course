package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware function to log request details
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		latency := time.Since(start)
		log.Printf("Log: %s %s %s %v", c.Request.Method, c.Request.URL.Path, c.ClientIP(), latency)
	}
}

func main() {
	router := gin.Default()

	// Use the Logger middleware
	router.Use(Logger())

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the home page!"})
	})

	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "About us page!"})
	})

	// Start server
	router.Run(":8080")
}
