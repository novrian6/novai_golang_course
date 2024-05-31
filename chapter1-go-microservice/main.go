package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL endpoint untuk menghitung jumlah karakter dalam sebuah string
	url := "http://localhost:8080/count-characters?text=Hello%20World"

	// Mengirim HTTP GET request ke layanan
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Membaca respons dari layanan
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan respons dari layanan
	fmt.Println("Response:", string(body))
}
