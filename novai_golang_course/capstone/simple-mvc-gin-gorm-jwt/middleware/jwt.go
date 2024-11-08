// /middleware/jwt.go
package middleware

import (
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthenticateJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("your-secret-key"), nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        c.Next()
    }
}