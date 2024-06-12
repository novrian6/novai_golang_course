// controllers/register_controller.go
package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController handles user registration
type RegisterController struct {
	DB *sql.DB
}

// NewRegisterController creates a new instance of RegisterController
func NewRegisterController(db *sql.DB) *RegisterController {
	return &RegisterController{DB: db}
}

// Register processes registration form submission
func (rc *RegisterController) Register(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Validate username and password here...

	// Insert new user into the database
	_, err = rc.DB.Exec("INSERT INTO users (username, email, password_hash,role) VALUES ($1, $2,$3,$4)", username, email, hashedPassword, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
