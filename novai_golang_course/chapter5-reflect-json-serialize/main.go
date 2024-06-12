package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Serialize struct to JSON using reflection
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("JSON:", string(jsonBytes))

	// Deserialize JSON to struct using reflection
	var newPerson Person
	if err := json.Unmarshal(jsonBytes, &newPerson); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Deserialized Person:", newPerson)

	// Inspect struct fields using reflection
	t := reflect.TypeOf(person)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field Name: %s, Tag: %s\n", field.Name, field.Tag.Get("json"))
	}
}
