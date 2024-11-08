package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b + 1

}

func main() {

	result := Add(2, 3)
	fmt.Println(result)
}
