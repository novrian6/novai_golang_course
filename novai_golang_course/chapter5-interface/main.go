package main

import "fmt"

// Shaper is an interface with a method to calculate the area
type Shaper interface {
    Area() float32
}

// Square struct implements the Shaper interface
type Square struct {
    side float32
}

// Area method calculates the area of the square
func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func main() {
    sq1 := new(Square)
    sq1.side = 5

    // Assigning sq1, which is of type *Square, to an interface variable of type Shaper
    var areaIntf Shaper = sq1

    // Calling the Area() method on the interface variable, which is implemented by Square
    fmt.Printf("The square has area: %f\n", areaIntf.Area())
}

