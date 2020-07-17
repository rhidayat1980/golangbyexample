package main

import (
	"fmt"
	"time"
)

func main() {

	// simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("")
	for i := 1; i <= 5; i++ {
		fmt.Println("hello")
	}

	// actually while in go is for loop
	fmt.Println("")
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// its similar with for loop below
	fmt.Println("")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// break in go
	fmt.Println("for break in go")
	i = 0
	for {
		fmt.Println(i)
		i++
		if i >= 10 {
			break
		}
	}

	// for continue in go

	fmt.Println("for continue in go")
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// nested for loop

	for i := 0; i < 3; i++ {
		fmt.Println("Outer loop iteration")
		for j := 0; j < 2; j++ {
			fmt.Printf("i = %d, j=%d", i, j)
		}
	}

	// for infinite loop
	fmt.Println("for infinite loop")
	i = 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Second * 1)
	}

}
