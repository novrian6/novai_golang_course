package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware to set secure headers
	r.Use(func(c *gin.Context) {
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Next()
	})

	// Route without authentication
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Simple authentication middleware example
	r.Use(func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken != "Bearer 123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	})

	// Route that requires authentication
	r.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the protected area!")
	})

	// Run the server with TLS
	//r.RunTLS(":8081", "server.crt", "server.key")

	// Run the server without TLS
	r.Run(":8081")
}
