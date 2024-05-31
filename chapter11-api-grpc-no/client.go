package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "chapter11-api-grpc-no/books.com"
)

const (
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewBookServiceClient(conn)

	// Test GetBooks
	getBooksResponse, err := client.GetBooks(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling GetBooks: %v", err)
	}
	log.Println("Books:")
	for _, book := range getBooksResponse.Books {
		log.Printf("ID: %d, Title: %s, Author: %s, Year: %d, ISBN: %s\n", book.Id, book.Title, book.Author, book.Year, book.Isbn)
	}

	// Test UpdateBook
	_, err = client.UpdateBook(context.Background(), &pb.Book{
		Id:     1,
		Title:  "Updated Title",
		Author: "Updated Author",
		Year:   2022,
		Isbn:   "987654321",
	})
	if err != nil {
		log.Fatalf("Error calling UpdateBook: %v", err)
	}
	log.Println("Book updated successfully")
}
