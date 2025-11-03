package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

//func hi() { 함수 이름이 소문자로 시작하면 외부에 노출되지 않는다
func Hi() {
	fmt.Println("Hi!")
}
