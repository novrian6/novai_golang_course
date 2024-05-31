package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func EnsureNotLoggedIn() gin.HandlerFunc {
	return func (c *gin.Context){


		loggedInInterface, _ := c.Get ("is_logged_in")

		loggedIn := loggedInInterface.(bool)
		if loggedIn{
			c.AbortWithStatus(http.StatusUnauthorized)

		}

	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			//if token, err := c.Cookie("token"); err != nil || token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}




func SetUserStatus() gin.HandlerFunc{
	return func(c *gin.Context){
		if token, err := c.Cookie ("token"); err==nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}