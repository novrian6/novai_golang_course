package main

import "fmt"

func Sum(nums ...int) int {

	total := 0

	for _, num := range nums {
		total += num

	}
	return total

}

func main() {

	total := Sum(1, 2, 3, 3, 4, 4, 2, 4, 1, 100)
	fmt.Println("total:" + fmt.Sprintf("%d", total))
}
