package main

import (
	"go-gin-app-me/handlers"
	middleware "go-gin-app-me/middleware"
)

func InitializeRoutes(){


	router.Use(middleware.SetUserStatus())



	router.GET ("/", handlers.ShowIndexPage)


	userRoutes := router.Group ("/u")
	{
		userRoutes.GET ("/login" , middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		userRoutes.POST ("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)

		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)
	}


}
