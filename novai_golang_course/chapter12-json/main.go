package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL JSON
	url := "https://api.github.com/users/golang"

	// Get response dari URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Baca body response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Parsing JSON
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Cetak nama dan public repos
	fmt.Println("Nama:", data["name"])
	fmt.Println("Public Repos:", data["public_repos"])
}
