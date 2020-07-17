package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("simple for loop")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	greeting := "hello"
	for i := 1; i <= 10; i++ {
		fmt.Println(strconv.Itoa(i) + "." + " " + greeting)
	}
}
