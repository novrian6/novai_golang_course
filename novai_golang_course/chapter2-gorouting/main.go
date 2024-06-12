package main

import (
	"fmt"
	"sync"
)

func goroutine2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("running in go routine 2")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("running in go routine 1")
	}()

	goroutine2(&wg)

	wg.Wait()

	fmt.Println("end")

}
