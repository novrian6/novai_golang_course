package main

import (
	"fmt"
)

// Define the custom error type
type MyError struct {
	Msg string
}

// Implement the error interface
func (e *MyError) Error() string {
	return e.Msg
}

// A function that returns an error
func doSomething() error {
	// Return an instance of MyError
	return &MyError{Msg: "something went wrong"}
}

func main() {
	// Call the function that may return an error
	if err := doSomething(); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("No error occurred.")
	}
}
