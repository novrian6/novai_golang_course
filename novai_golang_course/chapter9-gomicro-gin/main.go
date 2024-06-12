package main

import (
	"net/http"

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

func main() {
	// Use Go Micro to build the web service
	service := web.NewService(
		web.Name("countcharacter"),
	)

	// Initialize Gin
	ginRouter := gin.New()

	// Register the handler for the web service
	countService := new(CountCharacterService)
	ginRouter.GET("/count-characters", countService.CountCharacters)

	// Set the Gin router as the handler
	service.Handle("/", ginRouter)

	// Run the web service
	if err := service.Run(); err != nil {
		panic(err)
	}
}
