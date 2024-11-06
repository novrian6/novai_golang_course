package main

import (
	"fmt"
)

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Code    int
	Message string
}

// Error returns the error message of CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Custom Error %d: %s", e.Code, e.Message)
}

// FunctionThatMightFail is a function that returns a custom error.
func FunctionThatMightFail(input int) error {
	if input < 0 {
		return &CustomError{Code: 400, Message: "Input cannot be negative"}
	}
	return nil
}

func main() {
	err := FunctionThatMightFail(-1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = FunctionThatMightFail(5)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
