package main

import (
	"fmt"
	"time"
)

func main() {
	// Main goroutine
	fmt.Println("Starting the main goroutine")

	// Start a new goroutine
	go func() {
		fmt.Println("Hello from the goroutine1")
	}()
	go func() {
		fmt.Println("Hello from the goroutine2")
	}()
	go func() {
		fmt.Println("Hello from the goroutine3")
	}()
	go func() {
		fmt.Println("Hello from the goroutine4")
	}()
	go func() {
		fmt.Println("Hello from the goroutine5")
	}()

	// Continue execution in the main goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting the main goroutine")
}
