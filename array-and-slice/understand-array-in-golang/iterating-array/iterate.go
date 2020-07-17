package main

import "fmt"

func main() {

	letters := [3]string{"a", "b", "c"}

	// using for loop
	fmt.Println("Using for loop")
	len := len(letters)
	for i := 0; i < len; i++ {
		fmt.Println(i, letters[i])
	}

	// using for range operator
	fmt.Println("Using for range operator")
	for i, letter := range letters {
		fmt.Printf("%d %s\n", i, letter)
	}
}
