package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "novasecret123"

// LoginController handles login related operations
type LoginController struct {
	DB *sql.DB
}

// NewLoginController creates a new instance of LoginController
func NewLoginController(db *sql.DB) *LoginController {
	return &LoginController{
		DB: db,
	}
}

// ShowLoginForm displays the login form
func (lc *LoginController) ShowLoginForm(c *gin.Context) {
	// Render the login form HTML template
	c.HTML(http.StatusOK, "login.html", nil)
}

// Login handles login form submission
func (lc *LoginController) LoginHardCode(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check username and password against database (this is just a placeholder)
	if username == "admin" && password == "admin123" {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}

func (lc *LoginController) LoginNonJWT(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Query the database to check if the username and password match
	var storedPassword string
	err := lc.DB.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// Username not found in the database
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		// Other database error occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error:" + err.Error()})
		return
	}

	// Compare the stored password with the provided password
	if password != storedPassword {
		// Passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Authentication successful
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (lc *LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Query the database to check if the username and password match
	var storedPassword string
	err := lc.DB.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// Username not found in the database
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}
		// Other database error occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error:" + err.Error()})
		return
	}

	// Compare the stored password with the provided password
	/*if password != storedPassword {
		// Passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}*/

	// Compare the stored hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		// Passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Authentication successful
	// Display successful login message

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()     // Token expiration time (24 hours)
	signedToken, err := token.SignedString([]byte(secretKey)) //
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", signedToken, 3600, "/", "", false, true)

	// Return the token in the response headers
	//c.Header("Authorization", signedToken)
	//c.JSON(http.StatusOK, gin.H{"message": "Login successful"})

	c.JSON(http.StatusOK, gin.H{"token": signedToken})

}
