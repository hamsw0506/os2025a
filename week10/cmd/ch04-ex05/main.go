package main

import (
	"fmt"
	"log"

	// "week10/pkg/keyboard"
	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("숫자 입력 : ")
	num, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.1f\n", num)
}
