package main

import (
	"fmt"
)

func main() {
	numbers := make([]int, 3, 5)

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Printf("numbers: %d\n", numbers)
	fmt.Printf("length: %d\n", len(numbers))
	fmt.Printf("capacity: %d\n", cap(numbers))

	// append number 4
}
