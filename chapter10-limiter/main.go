package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(2, 5) // Allow 2 requests per 5 seconds

func main() {
	r := gin.Default()

	r.GET("/api/resource", handleResource)

	r.Run(":8080")
}

func handleResource(c *gin.Context) {
	// Check if the request exceeds the rate limit
	if limiter.Allow() {
		// If allowed, proceed with the request
		c.JSON(http.StatusOK, gin.H{
			"message": "Resource accessed successfully",
		})
	} else {
		// If rate limit exceeded, return a 429 Too Many Requests status code
		c.JSON(http.StatusTooManyRequests, gin.H{
			"message": "Rate limit exceeded. Please try again later.",
		})
		c.Abort() // Abort further processing of the request
	}
}

