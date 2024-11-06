package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// Connect to the PostgreSQL database
	db, err := gorm.Open(
		"postgres",
		"host=143.198.218.9 port=5432 user=postgres "+
			"dbname=OnlineStore password=P@ssw0rd sslmode=disable",
	)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Create a new user
	user := User{Name: "Nov Nov", Email: "nov@example.com"}
	db.Create(&user)

	// Retrieve all users
	var users []User
	db.Find(&users)
	fmt.Println("Get Users:", users)

	// Update user information
	db.Model(&user).Update("Name", "Jane Doe")
	fmt.Println("User is Updated:", user)

	// Delete user
	db.Delete(&user)

	// Count the total number of users
	var count int
	db.Model(&User{}).Count(&count)
	fmt.Println("Total users after delete:", count)
}
