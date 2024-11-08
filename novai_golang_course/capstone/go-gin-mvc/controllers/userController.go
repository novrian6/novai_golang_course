// /controllers/userController.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin-mvc/config"
	"go-gin-mvc/models"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
