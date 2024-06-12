package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-app-me/config"
	"go-gin-app-me/controller"
	//"go-gin-app-me/handlers"

	"gorm.io/gorm"
	"log"

)

var router *gin.Engine




var (
	db *gorm.DB = config.SetupDatabaseConnection()

)
var(
	authController controller.AuthController = controller.NewAuthController()

)

func main (){

	fmt.Println("hello C")

	defer config.CloseDatabaseConnection(db)



	router = gin.Default()

	router.GET ("/", controller.DisplayAll)

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login",authController.Login)
		authRoutes.POST("/register",authController.Register)


	}
	userRoutes := router.Group("/api/users")
	{

		userRoutes.POST("/all",controller.DisplayAll)


	}

	//router.LoadHTMLGlob("templates/*")


	//InitializeRoutes()

	err := router.Run(":8081")
	if err != nil {
		log.Fatalln(err)
	}

}



