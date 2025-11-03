package main

import (
	"fmt"
)

func main() {
	arrayBool := [3]bool{true, false, true}
	arrayInt := [2]int{-9, 11}
	for i := 0; i < len(arrayBool); i++ {
		fmt.Println(i, arrayBool[i])
		//fmt.Println(i, arrayInt[i])	//runtime error
	}
	fmt.Printf("%#v\n", arrayInt)
	fmt.Printf("%#v\n", arrayBool)
}
