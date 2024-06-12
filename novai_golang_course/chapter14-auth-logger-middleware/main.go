//test with reqbin with custom header

package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware global
	r.Use(LoggerMiddleware)

	// Route dengan middleware
	r.GET("/user/:id", AuthMiddleware, func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{"user_id": id})
	})

	r.Run(":8082")
}

// LoggerMiddleware adalah contoh middleware untuk logging
func LoggerMiddleware(c *gin.Context) {
	start := time.Now()

	// Next() memanggil handler berikutnya dalam antrian
	c.Next()

	// Logging setelah handler selesai dieksekusi
	end := time.Now()
	latency := end.Sub(start)
	fmt.Printf("Logger:%s - %s - %v\n", c.Request.Method, c.Request.URL.Path, latency)
}

// AuthMiddleware adalah contoh middleware untuk autentikasi
func AuthMiddleware(c *gin.Context) {
	// Contoh validasi autentikasi sederhana
	token := c.GetHeader("Authorization")
	if token != "secret_token" {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	// Lanjut ke handler berikutnya jika autentikasi berhasil
	c.Next()
}
