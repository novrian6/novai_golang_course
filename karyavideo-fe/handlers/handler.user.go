package handlers
/*
import (
	"github.com/gin-gonic/gin"
	m "go-gin-app-me/middleware"

//	models "go-gin-app-me/models"
	"math/rand"
	"net/http"
	"strconv"


)

func ShowLoginPage( c*gin.Context){
	m.Render(c, gin.H{
		"title" : "Login"}, "login.html")
}

func PerformLogin( c*gin.Context){
	username := c.PostForm("username")
	password:= c.PostForm("password")

	if models.IsUserValid(username, password){
		token := generateSessionToken()
		c.SetCookie("token",token,3600,"","",false,true)
		c.Set("is_logged_in",true)
		m.Render(c,gin.H{
			"title": "Successful Login"},
			"login-successful.html")
	}else{
		m.Render(c,gin.H{
			"title": "Fail Login"},
			"login-fail.html")
	}

}

func generateSessionToken() string{
	return strconv.FormatInt(rand.Int63(), 16)
}

func Logout(c *gin.Context) {

	//var sameSiteCookie http.SameSite;

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

*/