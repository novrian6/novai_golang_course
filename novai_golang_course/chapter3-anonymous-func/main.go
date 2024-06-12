package main

import (
	"fmt"
)

func main() {

	sum := func(a, b int) int { return a + b }(1, 2)
	fmt.Print("total:", sum)
}
