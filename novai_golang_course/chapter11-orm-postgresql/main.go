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
	// Koneksi ke database PostgreSQL
	db, err := gorm.Open("postgres", "host=rosie.db.elephantsql.com port=5432 user=bxsexxue dbname=bxsexxue password=5C758n0ZGuR11JKW8CoZEoqjtaJn_Mzm sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Migrasi model ke database (membuat tabel jika belum ada)
	db.AutoMigrate(&User{})

	// Contoh operasi CRUD
	// Insert data
	user := User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)

	// Query data
	var users []User
	db.Find(&users)
	fmt.Println("Users:", users)

	// Update data
	db.Model(&user).Update("Name", "Jane Doe")

	// Delete data
	db.Delete(&user)

	// Query data lagi setelah dihapus
	var count int
	db.Model(&User{}).Count(&count)
	fmt.Println("Total users after delete:", count)
}
