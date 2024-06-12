package main

import (
    "fmt"
)

func sum(arr []int, ch chan int) {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    ch <- sum // Mengirim hasil penjumlahan ke dalam channel
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    ch := make(chan int)

    go sum(nums[:len(nums)/2], ch) // Menghitung penjumlahan paruh pertama
    go sum(nums[len(nums)/2:], ch) // Menghitung penjumlahan paruh kedua

    partialSum1 := <-ch // Menerima hasil penjumlahan dari goroutine pertama
    partialSum2 := <-ch // Menerima hasil penjumlahan dari goroutine kedua

    totalSum := partialSum1 + partialSum2 // Menjumlahkan hasil penjumlahan paruh pertama dan kedua
    fmt.Println("Total sum:", totalSum)
}

