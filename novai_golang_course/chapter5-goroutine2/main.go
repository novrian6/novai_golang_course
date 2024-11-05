package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Message %d", i)
		fmt.Println("Send :", i)
		ch <- msg // Send message to the channel
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Close the channel after sending all messages
}

func receiver(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}

func main() {
	// Create a buffered channel of strings
	ch := make(chan string, 3)

	// Start sender goroutine
	go sender(ch)

	// Start receiver goroutine
	go receiver(ch)

	// Wait for a few seconds to allow goroutines to finish
	time.Sleep(3 * time.Second)
}
