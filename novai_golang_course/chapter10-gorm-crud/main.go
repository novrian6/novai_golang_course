// main.go
package main

import (
	"crud_gorm/config"
	"crud_gorm/models"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Migrate the schema (create the table if it doesn't exist)
	config.DB.AutoMigrate(&models.Book{})

	// Perform CRUD operations
	models.CreateBook("The Catcher in the Rye", "J.D. Salinger", 1951, "978-0-316-76948-0")
	models.GetBooks()
	models.GetBookByID(1)
	models.UpdateBook(1, "The Catcher in the Rye - Updated")
	models.DeleteBook(1)
}
