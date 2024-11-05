package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	//fmt.Print("Hello, world!")
	for index, value := range numbers {
		fmt.Printf("Index %d, Value:%d\n", index, value)
	}
}
