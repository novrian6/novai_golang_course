package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	// Define a handler for the /foo route
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %q", html.EscapeString("/foo"))
	})

	// Define a handler for the /bar route
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %q", html.EscapeString("/bar"))
	})

	// Start serving the application on port 8080
	http.ListenAndServe(":8081", nil)
}
