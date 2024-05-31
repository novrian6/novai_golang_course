package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//s"document_platform/internal/routes"

	"com.hypnovai.documentpublishing/internal/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "172.16.137.133"
	port     = 5432
	user     = "postgres"
	password = "zuruck"
	dbname   = "document_platform_db"
)

func main() {
	// Connect to PostgreSQL database
	db := connectDB()

	// Initialize Gin
	router := gin.Default()

	// Get the absolute path of the directory containing the main.go file
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Navigate up one level to the project root
	projectRoot := filepath.Dir(filepath.Dir(dir))

	// Define the relative path to the views directory
	viewsDir := filepath.Join(projectRoot, "internal", "views")

	router.LoadHTMLGlob(filepath.Join(viewsDir, "*.html"))

	// Load routes
	routes.SetupRoutes(router, db)

	// Start the Gin server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func connectDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	return db
}
