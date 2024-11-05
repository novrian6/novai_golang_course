package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "my_secret_key"

func main() {
	r := gin.Default()

	// Route for generating tokens
	r.POST("/login", handleLogin)

	// Middleware to check JWT on every request
	r.Use(authMiddleware())

	// Protected routes
	r.GET("/api/general", handleGeneralResource)
	r.GET("/api/admin", adminMiddleware(), handleAdminResource)

	r.Run(":8080")
}

func handleLogin(c *gin.Context) {
	// In a real application, authenticate the user (this is just an example)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check user credentials
	var role string
	if username == "admin" && password == "password" {
		role = "admin"
	} else if username == "user" && password == "password" {
		role = "user"
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop further processing if unauthorized
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next() // Proceed to the next handler if authorized
	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func handleGeneralResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "General resource accessed successfully",
	})
}

func handleAdminResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Admin resource accessed successfully",
	})
}
