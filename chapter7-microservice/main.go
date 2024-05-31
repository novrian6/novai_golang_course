package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Inisialisasi Gin
    router := gin.Default()

    // Endpoint untuk menghitung jumlah karakter dalam sebuah string
    router.GET("/count-characters", func(c *gin.Context) {
        // Ambil query parameter "text" dari URL
        text := c.Query("text")

        // Hitung jumlah karakter dalam string
        count := len(text)

        // Kirim hasil hitungan sebagai respons
        c.JSON(http.StatusOK, gin.H{
            "text":  text,
            "count": count,
        })
    })

    // Menjalankan server di port 8080
    router.Run(":8082")
}

