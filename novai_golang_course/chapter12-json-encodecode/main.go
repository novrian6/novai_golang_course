package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct for the data we want to encode/decode
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func main() {
	// Create an instance of the struct
	person := Person{
		Name:   "John Doe",
		Age:    30,
		Email:  "john.doe@example.com",
		Active: true,
	}

	// Encode the struct into JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error encoding JSON: %s", err)
	}
	fmt.Println("Encoded JSON:", string(jsonData))

	// Define a JSON string
	jsonString := `{"name":"Jane Smith","age":25,"email":"jane.smith@example.com","active":false}`

	// Decode the JSON string into a struct
	var person2 Person
	err = json.Unmarshal([]byte(jsonString), &person2)
	if err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}
	fmt.Printf("Decoded Struct: %+v\n", person2)
}
