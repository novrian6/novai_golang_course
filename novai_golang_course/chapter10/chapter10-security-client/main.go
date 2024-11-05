package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Buat request dengan token
	req, err := http.NewRequest("GET", "http://localhost:8080/api/resource", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set header Authorization dengan token
	req.Header.Set("Authorization", "my_secret_token")

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Baca response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

