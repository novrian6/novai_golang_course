// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Welcome to our website</h1>")
	fmt.Fprintln(w, `<a href="/login">Go to Login</a>`)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "user" && password == "pass" {
			fmt.Fprintln(w, "<h1>Login Successful</h1>")
		} else {
			fmt.Fprintln(w, "<h1>Login Failed</h1>")
		}
		return
	}

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
