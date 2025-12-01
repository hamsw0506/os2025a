package main

import "fmt"

func main() {
	fmt.Println("1. 프로그램 시작")

	panic("심각한 에러 발생")

	fmt.Println("해당 줄은 실행 안됨")
}
