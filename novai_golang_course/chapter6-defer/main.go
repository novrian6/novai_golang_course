package main

import "fmt"

func main() {
    defer fmt.Println("World") // Deferring this function call until the end of the enclosing function (main)

    fmt.Print("Hello, ")

    // Multiple defer statements are pushed onto a stack and executed in last-in-first-out order
    defer fmt.Println("Deferred 1")
    defer fmt.Println("Deferred 2")
    defer fmt.Println("Deferred 3")

    fmt.Println("End of main function")
}

