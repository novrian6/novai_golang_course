package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter int
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	// Allow time for all goroutines to finish
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Final Counter Value (with synchronization):", counter)
}
