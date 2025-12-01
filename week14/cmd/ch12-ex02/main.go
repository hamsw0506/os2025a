package main

import "fmt"

func main() {
	defer fmt.Println("프로그램 종료 - 정리 작업")
	fmt.Println("1. 프로그램 시작")

	panic("심각한 에러 발생")

	fmt.Println("해당 줄은 실행 안됨")
}
