package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	
)

var router *gin.Engine


const (
	DbHost = "db"
	DbUser = "nova"
	DbPassword= "password"
	DbName = "dev"
	Migration = `CREATE TABLE IF NOT EXISTS bulletins (
	id serial PRIMARY KEY,
	author text NOT NULL,
	content text NOT NULL,
	created_at timestamp with time zone DEFAULT current_timestamp
	)`

)

var db *sql.DB

func main (){

	fmt.Println("hello C")

/*	db, err := sql.Open("mysql", "root:password123@tcp(db:3306)/testdb")

	if err != nil {
		panic (err)
	}

	defer db.Close()*/

	router = gin.Default()

	router.LoadHTMLGlob("templates/*")


	InitializeRoutes()

	err := router.Run(":8083")
	if err != nil {
		log.Fatalln(err)
	}

}



