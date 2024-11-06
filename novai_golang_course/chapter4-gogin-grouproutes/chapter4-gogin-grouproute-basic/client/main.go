package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Consume API v1
	responseV1, err := http.Get("http://localhost:8080/api/v1/hello")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer responseV1.Body.Close()

	bodyV1, err := ioutil.ReadAll(responseV1.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response from API v1:", string(bodyV1))

	// Consume API v2
	responseV2, err := http.Get("http://localhost:8080/api/v2/hello")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer responseV2.Body.Close()

	bodyV2, err := ioutil.ReadAll(responseV2.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response from API v2:", string(bodyV2))
}
