package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)


func SetupDatabaseConnection() *gorm.DB{
	err:= godotenv.Load()
	if err != nil {
		panic ("Failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser, dbPass,dbHost,dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic ("Failed to connect db")

	}
	return db

}

func CloseDatabaseConnection(db *gorm.DB){
	dbSQL, err :=db.DB()
	if err != nil {
		panic ("failed to close connection")
	}
	dbSQL.Close()
}