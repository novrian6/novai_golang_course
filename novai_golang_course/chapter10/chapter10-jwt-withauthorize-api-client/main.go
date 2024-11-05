package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Define command line flags for username and password
	username := flag.String("username", "", "Username for login")
	password := flag.String("password", "", "Password for login")
	flag.Parse()

	// Validate input
	if *username == "" || *password == "" {
		fmt.Println("Username and password are required")
		fmt.Println("Syntax: go run main.go -username=[user] -password=[pasword]")

		return
	}

	// test  Log in ass user to get a token
	token, err := login(*username, *password)

	// test  Log in ass admin to get a token
	//token, err := login("admin", "password")

	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}

	// Access general resource
	fmt.Println("Accessing general resource:")
	accessResource(token, "http://localhost:8080/api/general")

	// Access admin resource
	fmt.Println("Accessing admin resource:")
	accessResource(token, "http://localhost:8080/api/admin")
}

func login(username, password string) (string, error) {
	// Create the form data
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	// Send the login request
	resp, err := http.PostForm("http://localhost:8080/login", formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response to get the token
	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result["token"], nil
}

func accessResource(token, url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Authorization header with the token
	req.Header.Set("Authorization", token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
