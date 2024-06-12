

package main

import (
	"fmt"
	"reflect"
)

func PrintFields(v interface{}) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() != reflect.Struct {
		fmt.Println("Expected a struct")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user := User{Name: "Novrian", Email: "novrian@example.com", Age: 30}
	PrintFields(user)
}

