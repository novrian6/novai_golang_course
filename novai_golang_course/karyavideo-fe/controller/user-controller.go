package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-app-me/config"
	"go-gin-app-me/models"
	"gorm.io/gorm"

	"net/http"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

)

func DisplayAll (c *gin.Context){


	//db:=c.MMustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK,gin.H{
		"data":users,
	})
}
