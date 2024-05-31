package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string `bson:"title"`
	Author string `bson:"author"`
	Year   int    `bson:"year"`
	ISBN   string `bson:"isbn"`
}

func main() {
	// Set up MongoDB connection string
	uri := "mongodb+srv://liebera6:Pa$$w0rd@cluster0.gxtuuq7.mongodb.net/"

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(uri)

	// Set connection context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Select database and collection
	collection := client.Database("bookstore_db").Collection("books")

	// Create slice to store book data
	var books []Book

	// Find all documents in collection
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Iterate through cursor results and append to books slice
	for cursor.Next(ctx) {
		var book Book
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	// Display book data
	for _, book := range books {
		fmt.Printf("Title: %s, Author: %s, Year: %d, ISBN: %s\n", book.Title, book.Author, book.Year, book.ISBN)
	}
}
