package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "rosie.db.elephantsql.com "
	port     = 5432
	user     = "bxsexxue"
	password = "5C758n0ZGuR11JKW8CoZEoqjtaJn_Mzm"
	dbname   = "bxsexxue"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	ISBN   string
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, title, author, year, isbn FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.ISBN)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d, ISBN: %s\n", book.ID, book.Title, book.Author, book.Year, book.ISBN)
	}
}
