package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	start := time.Now()

	// Before executing the main handler
	log.Printf("Started %s %s", c.Request.Method, c.Request.RequestURI)

	// Pass on to the next middleware/handler
	c.Next()

	// After executing the main handler
	log.Printf("Completed %s in %v", c.Request.RequestURI, time.Since(start))
}

func main() {
	r := gin.Default()

	// Use the LoggingMiddleware for all routes
	r.Use(LoggingMiddleware)

	// Define a simple JSON response endpoint
	r.GET("/json", func(c *gin.Context) {
		data := map[string]string{
			"message": "Hello, JSON!",
		}
		c.JSON(http.StatusOK, data)
	})

	// Start the server
	r.Run(":8080")
}
