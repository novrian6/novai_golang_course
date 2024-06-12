package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

// LoginForm represents the structure of the login form data
type LoginForm struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}
// validateInput validates the login form data
func validateInput(form LoginForm) error {
	validate := validator.New()
	return validate.Struct(form)
}
// sanitizeInput sanitizes the login form data
func sanitizeInput(form *LoginForm) {
	p := bluemonday.StrictPolicy()
	form.Username = p.Sanitize(form.Username)
	form.Password = p.Sanitize(form.Password)
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm

		// Bind the JSON data from the request body to the LoginForm struct
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Sanitize the input
		sanitizeInput(&form)

		// Validate the input
		if err := validateInput(form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Dummy validation for username and password
		if form.Username == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		}
	})

	// Run the server
	r.Run(":8080")
}
