package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// Define a rate limiter with 1 request per second and a burst of 5.
var limiter = rate.NewLimiter(1, 5)

// Middleware to check the rate limit.
func rateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		c.Abort()
		return
	}
	c.Next()
}

func main() {
	r := gin.Default()

	// Apply the rate limiting middleware
	r.Use(rateLimiter)

	// Define a simple route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the rate-limited endpoint!")
	})

	// Run the server
	r.Run(":8080")
}
