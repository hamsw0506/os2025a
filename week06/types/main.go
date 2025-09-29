package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name = "Ham Sunwoo"
	// fmt.Println(name)
	// fmt.Println(math.Ceil(2.71))
	// fmt.Println(strings.Title("head first go"))

	// var name string
	// name = "Ham Sunwoo"
	// fmt.Println(name)

	var name float64
	name = 2.71
	fmt.Println(name, reflect.TypeOf(name))
}
