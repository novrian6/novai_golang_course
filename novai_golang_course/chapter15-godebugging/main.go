package main

import "fmt"

func main() {
	fmt.Println("Starting the program")
	result := Add(2, 3)
	fmt.Println("Result:", result)
}

func Add(a, b int) int {
	hasil := a - b
	return hasil // Ini akan menghasilkan kesalahan
}
