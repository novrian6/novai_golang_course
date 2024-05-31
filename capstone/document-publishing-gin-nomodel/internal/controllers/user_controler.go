// controllers/user_controller.go
package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserController handles user-related actions
type UserController struct {
	DB *sql.DB
}

// NewUserController creates a new instance of UserController
func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

// GetUserList retrieves a list of all users from the database
func (uc *UserController) GetUserList(c *gin.Context) {
	// Check if user is authenticated
	if _, exists := c.Get("claims"); !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Query the database to retrieve all users' names and emails
	rows, err := uc.DB.Query("SELECT username, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}
	defer rows.Close()

	var users []gin.H
	for rows.Next() {
		var name, email string
		err := rows.Scan(&name, &email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user data"})
			return
		}
		users = append(users, gin.H{"name": name, "email": email})
	}

	c.JSON(http.StatusOK, users)
	//c.HTML(http.StatusOK, "users.html", users)

}

// ChangePassword updates the user's password in the database
func (uc *UserController) ChangePassword(c *gin.Context) {
	// Retrieve user ID and new password from request
	userName := c.PostForm("username")
	newPassword := c.PostForm("new_password")

	// Authenticate and authorize user
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	claimsUserName := claims.(jwt.MapClaims)["username"].(string)
	// Check if the user ID from the JWT token matches the requested user ID
	if userName != claimsUserName {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Perform validation (e.g., check if the new password meets requirements)

	// Update the user's password in the database
	_, err = uc.DB.Exec("UPDATE users SET password_hash = $1 WHERE username = $2", hashedPassword, userName)
	if err != nil {
		log.Println("Error executing SQL query:", err)
		log.Println("SQL query:")
		log.Println("New password:", newPassword)
		log.Println("Username:", userName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
