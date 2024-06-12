package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host     = "rosie.db.elephantsql.com"
	port     = 5432
	user     = "bxsexxue"
	password = "5C758n0ZGuR11JKW8CoZEoqjtaJn_Mzm"
	dbname   = "bxsexxue"
	secret   = "the_secret"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	ISBN   string `json:"isbn"`
}

type BookService struct{}

func (s *BookService) GetBooks(ctx context.Context, req *GetBooksRequest) (*GetBooksResponse, error) {
	// Implement fetching books from the database
	books, err := fetchBooksFromDB()
	if err != nil {
		return nil, err
	}

	return &GetBooksResponse{Books: books}, nil
}

func (s *BookService) UpdateBook(ctx context.Context, req *UpdateBookRequest) (*UpdateBookResponse, error) {
	// Implement updating a book in the database
	err := updateBookInDB(req.Book)
	if err != nil {
		return nil, err
	}

	return &UpdateBookResponse{}, nil
}

func fetchBooksFromDB() ([]*Book, error) {
	// Database connection setup
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query database for books
	rows, err := db.Query("SELECT id, title, author, year, isbn FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process rows and create Book objects
	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.ISBN)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func updateBookInDB(book *Book) error {
	// Database connection setup
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Update book in the database
	_, err = db.Exec("UPDATE books SET title=$1, author=$2, year=$3, isbn=$4 WHERE id=$5",
		book.Title, book.Author, book.Year, book.ISBN, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()
	RegisterBookServiceServer(server, &BookService{})
	log.Println("gRPC server started on port :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
