package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginForm represents the structure of the login form data
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	CSRF     string `json:"csrf"`
}

func main() {
	r := gin.Default()

	// CSRF protection middleware
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "POST" {
			token := c.PostForm("csrf")
			if token != "valid_csrf_token" {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid CSRF token"})
				return
			}
		}
		c.Next()
	})

	// Login route
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm

		// Bind the JSON data from the request body to the LoginForm struct
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the password is equal to "P@ssw0rd"
		if form.Password == "P@ssw0rd" {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password"})
		}
	})

	// Run the server
	r.Run(":8080")
}
