package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 3.2
	var width int = 2
	fmt.Println("Area is", int(length)*width)
	fmt.Println("length > width?", int(length) > width)
	fmt.Println("con", reflect.TypeOf(int(length)))
	fmt.Println("origin", reflect.TypeOf(length))
}
