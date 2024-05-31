package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL untuk melakukan permintaan HTTP
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Melakukan permintaan HTTP GET
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Membaca isi respons
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mencetak respons
	fmt.Println(string(body))
}

