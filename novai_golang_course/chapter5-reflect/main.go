package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // Outputs the type of x

	v := reflect.ValueOf(x) // Gets the reflection Value of x
	fmt.Println("value:", v) // Outputs the value of v, which is a reflection of x

	fmt.Println("type:", v.Type()) // Outputs the type of v, should be same as x
	fmt.Println("kind:", v.Kind()) // Outputs the kind of v, which is a specific kind of the type
	fmt.Println("value:", v.Float()) // Retrieves the float value of v

	fmt.Println(v.Interface()) // Converts v back to an interface and then to its original type
	fmt.Printf("value is %5.2e\n", v.Interface()) // Formatted output of the value

	y := v.Interface().(float64) // Type asserts the interface value back to float64
	fmt.Println(y) // Outputs the original value of x
}

