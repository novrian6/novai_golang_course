// main.go
package main

import (
	"go-gin-mvc/routes"

	"go-gin-mvc/config"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080") // Run on port 8080
}
