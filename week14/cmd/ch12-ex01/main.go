package main

import (
	"fmt"
)

func main() {
	//Last in First Out, 스택 구조
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("Main logic")
}
