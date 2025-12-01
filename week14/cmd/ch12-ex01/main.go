package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("파일 오픈 실패 : ", err)
		return
	}

	defer file.Close()

	fmt.Println("파일 오픈 실행")
}
