// main.go
package main

import (
	"./routes"

	"simple-mvc-gin-gorm-jwt/config"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080") // Run on port 8080
}
