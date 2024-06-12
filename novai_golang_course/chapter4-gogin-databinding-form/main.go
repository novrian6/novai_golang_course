package main
import (
 "net/http"
 "github.com/gin-gonic/gin"
)

type User struct {
 Name  string `form:"name" binding:"required"`
 Email string `form:"email" binding:"required,email"`
}

func main() {
 r := gin.Default()
 // POST /user endpoint
 r.POST("/user", func(c *gin.Context) {
  var user User
  // Bind the form data from the request body to the User struct
  if err := c.ShouldBind(&user); err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
   return
  }
  // If data binding is successful, return the user information
  c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
 })
 // Run the server
 r.Run(":8080")
}
