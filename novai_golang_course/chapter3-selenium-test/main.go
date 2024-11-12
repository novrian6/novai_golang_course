// main.go
package main

import (
	"fmt"
	"net/http"
)

// LoginHandler handles the login request
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Simple authentication logic
		if username == "user" && password == "pass" {
			fmt.Fprintln(w, "<h1>Login Successful</h1>")
		} else {
			fmt.Fprintln(w, "<h1>Login Failed</h1>")
		}
		return
	}

	// Display the login form if the method is not POST
	tmpl := `<h1>Login</h1>
	<form method="POST" action="/login">
		<label>Username:</label>
		<input type="text" name="username"><br>
		<label>Password:</label>
		<input type="password" name="password"><br>
		<input type="submit" value="Login">
	</form>`
	w.Write([]byte(tmpl))
}

// HomeHandler handles the home page request
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome to our website</h1>")
	fmt.Fprintln(w, `<a href="/login">Go to Login</a>`)
}

// main function starts the web server
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
