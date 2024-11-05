package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// Log in to get a token
	token, err := login("admin", "password")
	if err != nil {
		fmt.Println("Error logging in:", err)
		return
	}

	// Create a request with the token
	req, err := http.NewRequest("GET", "http://localhost:8080/api/resource", nil)
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

func login(username, password string) (string, error) {
	// Create the login payload
	///data := map[string]string{"username": username, "password": password}
	///jsonData, err := json.Marshal(data)
	///if err != nil {
	///	return "", err
	///}

	// Send the login request
	///resp, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(jsonData))

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
