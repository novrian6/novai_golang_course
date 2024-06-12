package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Route for GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Route for POST /submit
	r.POST("/submit", func(c *gin.Context) {
		c.String(http.StatusOK, "Submission successful")
	})

	// Route for GET /user/:name
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, fmt.Sprintf("Hello, %s", name))
	})

	// Route for GET /welcome with query strings first_name and last_name
	r.GET("/welcome", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.Query("last_name")
		c.String(http.StatusOK, fmt.Sprintf("Welcome, %s %s", firstName, lastName))
	})

	// Route for POST /form_post with input text named "message"
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		c.String(http.StatusOK, fmt.Sprintf("Received message: %s", message))
	})

	// Route for POST /upload to handle file upload
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// Run the server
	r.Run(":8080")
}
