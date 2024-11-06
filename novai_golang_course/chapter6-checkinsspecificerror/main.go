package main

import (
	"errors"
	"fmt"
)

// Define a target error to compare against
var targetError = errors.New("a specific error occurred")

// A function that returns an error, possibly wrapping the target error
func doSomething() error {
	// Wrapping the target error with additional context
	return fmt.Errorf("an error occurred: %w", targetError)
}

func main() {
	// Call the function that may return an error
	err := doSomething()

	// Check if the error matches the target error
	if errors.Is(err, targetError) {
		// Handle specific error
		fmt.Println("Handling the specific error:", targetError)
	} else {
		fmt.Println("An unexpected error occurred:", err)
	}
}
