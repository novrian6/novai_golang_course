package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

// CountCharacterService is the implementation of a service that counts the number of characters in a string
type CountCharacterService struct{}

// CountCharacters is the method to count the number of characters in a string
func (s *CountCharacterService) CountCharacters(c *gin.Context) {
	// Get the query parameter "text" from the URL
	text := c.Query("text")

	// Count the number of characters in the string
	count := len(text)

	// Send the count as the response
	c.JSON(http.StatusOK, gin.H{
		"text":  text,
		"count": count,
	})
}

// ToLowerCharacterService is the implementation of a service that converts string to lowercase
type ToLowerCharacterService struct{}

// ToLowerCharacters is the method to convert string to lowercase
func (s *ToLowerCharacterService) ToLowerCharacters(c *gin.Context) {
	// Get the query parameter "text" from the URL
	text := c.Query("text")

	// Convert the string to lowercase
	lowerText := strings.ToLower(text)

	// Send the lowercase string as the response
	c.JSON(http.StatusOK, gin.H{
		"text": lowerText,
	})
}

func main() {
	// Use Go Micro to build the web service
	service := web.NewService(
		web.Name("countcharacter"),
	)

	// Initialize Gin
	ginRouter := gin.New()

	// Register the handlers
	countService := new(CountCharacterService)
	lowerService := new(ToLowerCharacterService)

	ginRouter.GET("/count-characters", countService.CountCharacters)
	ginRouter.GET("/to-lower", lowerService.ToLowerCharacters)

	// Set the Gin router as the handler
	service.Handle("/", ginRouter)

	// Run the web service
	if err := service.Run(); err != nil {
		panic(err)
	}
}
