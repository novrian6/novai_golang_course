package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Middleware untuk mengecek token setiap request
	r.Use(authMiddleware())

	r.GET("/api/resource", handleResource)

	r.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		// Cek apakah token valid (contoh sederhana)
		if token != "my_secret_token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Menghentikan proses lebih lanjut jika tidak authorized
			return
		}

		c.Next() // Lanjut ke handler berikutnya jika authorized
	}
}

func handleResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Resource accessed successfully",
	})
}

