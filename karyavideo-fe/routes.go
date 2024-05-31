package main
/*

import (
	"go-gin-app-me/controller"
	"go-gin-app-me/handlers"
	middleware "go-gin-app-me/middleware"
)

var(
	authController controller.AuthController = controller.NewAuthController()
)
func InitializeRoutes(){


	router.Use(middleware.SetUserStatus())



	router.GET ("/", handlers.ShowIndexPage)

 	authRoutes := router.Group("api/auth")
 	{
 		authRoutes.POST("/login",authController.Login())
		authRoutes.POST("/register",authController.Register())

	}

	userRoutes := router.Group ("/u")
	{
		userRoutes.GET ("/login" , middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		userRoutes.POST ("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

		userRoutes.POST ("/getallusers", middleware.EnsureNotLoggedIn(), handlers.GetAllUsers)

		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)
	}


}
*/