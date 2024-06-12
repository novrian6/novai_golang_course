package main

import "fmt"

// Function to simulate a division operation
func divide(a, b int) int {
	// Recover from panic if it occurs
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Check if divisor is zero
	if b == 0 {
		panic("division by zero")
	}

	// Perform the division operation
	return a / b
}

func main() {
	// Example 1: Normal division operation
	result := divide(10, 2)
	fmt.Println("Result of 10 / 2:", result)

	// Example 2: Division by zero (will cause panic)
	result = divide(10, 0)                   // This will cause panic
	fmt.Println("Result of 10 / 0:", result) // This line won't be executed
}
