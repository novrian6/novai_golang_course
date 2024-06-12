package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "chapter11-api-grpc-no/books.com"
)

const (
	host     = "rosie.db.elephantsql.com"
	port     = 5432
	user     = "bxsexxue"
	password = "5C758n0ZGuR11JKW8CoZEoqjtaJn_Mzm"
	dbname   = "bxsexxue"
)

var db *sql.DB

type server struct {
	pb.UnimplementedBookServiceServer
}

func main() {
	// Set up database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Set up gRPC server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterBookServiceServer(srv, &server{})
	log.Println("gRPC server running on port :8080")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) GetBooks(ctx context.Context, req *pb.Empty) (*pb.BookList, error) {
	rows, err := db.Query("SELECT id, title, author, year, isbn FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*pb.Book
	for rows.Next() {
		var book pb.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year, &book.Isbn)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.BookList{Books: books}, nil
}

func (s *server) UpdateBook(ctx context.Context, book *pb.Book) (*pb.Empty, error) {
	_, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3, isbn=$4 WHERE id=$5",
		book.Title, book.Author, book.Year, book.Isbn, book.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
