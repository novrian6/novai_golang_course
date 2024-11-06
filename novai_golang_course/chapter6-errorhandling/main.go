package main

import (
	"errors"
	"fmt"
)

// Define the doSomething function
func doSomething() error {
	// Simulate an error
	return errors.New("something went wrong")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("No error occurred.")
	}
}
