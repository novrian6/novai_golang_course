package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	for i := 0; i < 100; i++ {
		go func() {
			counter++
		}()
	}

	// Allow time for all goroutines to finish
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Final Counter Value (without synchronization):", counter)
}
