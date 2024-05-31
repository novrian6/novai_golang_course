package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the request information
		duration := time.Since(start)
		fmt.Printf("Log: [%s] %s %s - %s\n", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Us")
}

func main() {
	r := mux.NewRouter()

	// Use the LoggingMiddleware for all routes
	r.Use(LoggingMiddleware)

	// Define the routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/about", AboutHandler).Methods("GET")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
