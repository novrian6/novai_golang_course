package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define the Submission struct
type Submission struct {
	FirstName string
	LastName  string
}

func main() {
	r := gin.Default()

	// Route to render the form
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	// Route to handle form submission
	r.POST("/submit", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")

		submission := Submission{
			FirstName: firstName,
			LastName:  lastName,
		}

		c.HTML(http.StatusOK, "result.html", submission)
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Start the server
	r.Run(":8085")
}
