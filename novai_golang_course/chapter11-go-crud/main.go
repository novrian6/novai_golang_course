package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
    // Database connection settings
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        "143.198.218.9", 5432, "postgres", "P@ssw0rd", "OnlineStore",
    )
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Execute a simple insert statement
    result, err := db.Exec(
        "INSERT INTO public.books (title, author, year, isbn) VALUES ($1, $2, $3, $4)",
        "The Catcher in the Rye", "J.D. Salinger", 1951, "978-0-316-76948-0",
    )
    if err != nil {
        log.Fatal("Failed to insert record:", err)
    }

    // Check the result of the insert operation
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Fatal("Failed to retrieve number of rows affected:", err)
    }

    fmt.Printf("Inserted record successfully, %d row(s) affected.\n", rowsAffected)
}