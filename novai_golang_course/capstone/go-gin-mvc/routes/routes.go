// /routes/routes.go
package routes

import (
	"go-gin-mvc/controllers"
	"go-gin-mvc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterUser)
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthenticateJWT())
	{
		// Define protected routes here
	}
	return r
}
