package main

import (
	"errors"
	"fmt"
)

func division(x int, y int) (float32, error) {

	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return float32(x / y), nil

}

func main() {

	result, err := division(10, 2)

	fmt.Println(result)
	fmt.Println(err)

}
