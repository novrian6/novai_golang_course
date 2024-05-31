package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Menggunakan mode Debug untuk memudahkan debugging
	r := gin.Default()

	// Menambahkan handler untuk route GET "/hello"
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Menjalankan server di port 8080
	r.Run(":8080")
}

