package handlers

import (
	"github.com/gin-gonic/gin"
	m "go-gin-app-me/middleware"
	models2 "go-gin-app-me/models"
)

func ShowIndexPage(c *gin.Context){
	articles := models2.GetAllArticles()
	m.Render(c,gin.H{
		"title": "HomePage",
		"payload": articles},"index.html")
}




