package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// Version A: Sequential execution, not using goroutines
	fmt.Println("Version A: Sequential execution")
	for ix := range values { // ix is the index
		func(ix int) {
			fmt.Print(ix, " ")
		}(ix) // Pass ix as an argument to the closure
	}
	fmt.Println()

	// Version B: Incorrect goroutine usage, might not print expected results due to the closure capturing the loop variable
	fmt.Println("Version B: Incorrect goroutine usage")
	for ix := range values {
		go func() {
			fmt.Print(ix, " ") // ix is captured by the closure, might lead to unexpected results
		}()
	}
	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println()

	// Version C: Correct goroutine usage, passing the loop variable as an argument to the closure
	fmt.Println("Version C: Correct goroutine usage")
	for ix := range values {
		go func(ix int) {
			fmt.Print(ix, " ")
		}(ix) // Pass ix as an argument to ensure each goroutine gets the correct value
	}
	time.Sleep(1 * time.Second) // Wait for goroutines to finish
	fmt.Println()
}

