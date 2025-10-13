package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var fmt string = "inha"
	// var int int = 7
	// fmt.Println(int)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) //repot the error and exit program
	}

	i = strings.TrimSpace(i)
	score, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(err)
	}

	var pf string
	if score >= 60 {
		pf = "Pass"
	} else {
		pf = "Fail"
	}
	fmt.Println(pf, score)
}
