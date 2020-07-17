package main

import "fmt"

func main() {

	letters := []string{"a", "b", "c", "d", "e", "f"}

	// using for loop
	fmt.Println("Using for loop")
	len := len(letters)
	for i := 0; i < len; i++ {
		fmt.Println(i, letters[i])
	}

	for k, v := range letters {
		fmt.Println(k, " = ", v)
	}
}
