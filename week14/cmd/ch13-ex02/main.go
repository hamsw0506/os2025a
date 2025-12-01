package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	x := <-ch
	fmt.Println(x)
}
