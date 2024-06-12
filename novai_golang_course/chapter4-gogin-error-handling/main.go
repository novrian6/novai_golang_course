package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Custom middleware for error handling
	r.Use(ErrorHandler())
	r.GET("/", func(c *gin.Context) {
		// Set content type to HTML
		c.Header("Content-Type", "text/html; charset=utf-8")
		// Write HTML content directly into the response
		c.String(http.StatusOK, "<h1>Main Page</h1>")
	})

	// Define a route that will intentionally throw an error
	r.GET("/divbyzero", func(c *gin.Context) {
		// Simulate an error by dividing by zero
		var x, y int
		z := x / y // This will cause a runtime error
		fmt.Println(z)
	})

	// Run the server
	fmt.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}

// ErrorHandler is a custom middleware for handling errors
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Println("Recovered from panic:", err)

				// Return an error response to the client
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Recovered Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
