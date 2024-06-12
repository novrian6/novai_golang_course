// main.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/register", showRegisterForm)
	r.POST("/register", register)
	r.GET("/login", showLoginForm)
	r.POST("/login", login)
	//	r.GET("/users", authRequired, listUsers)
	r.GET("/users", authRequired, listUsers)

	r.Run(":8081")
}

func showRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	resp, err := http.Post("http://localhost:8080/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to register"})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func showLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	resp, err := http.PostForm("http://localhost:8080/api/login", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to login"})
		return
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	token := result["token"]

	session := sessions.Default(c)
	session.Set("token", token)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func listUsers(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token")
	if token == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/api/users", nil)
	req.Header.Set("Authorization", token.(string))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Failed to fetch users"})
		return
	}

	var users []User
	json.NewDecoder(resp.Body).Decode(&users)

	c.HTML(http.StatusOK, "users.html", gin.H{"users": users})
}

func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	token := session.Get("token")
	if token == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Next()
}
