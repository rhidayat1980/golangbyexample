package main

import "fmt"

func main() {

	numbers := [5]int{1, 2, 3, 4, 5}

	// both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1 = %v\n", num1)
	fmt.Printf("length = %d\n", len(num1))
	fmt.Printf("capacity = %d\n", cap(num1))

	// only start
	num2 := numbers[2:]
	fmt.Printf("\nonly start")
	fmt.Printf("num2 = %v\n", num2)
	fmt.Printf("length = %d\n", len(num2))
	fmt.Printf("capacity = %d\n", cap(num2))

	// only end
	num3 := numbers[:3]
	fmt.Println("only end")
	fmt.Printf("num3 = %v\n", num3)
	fmt.Printf("length = %d\n", len(num3))
	fmt.Printf("capacity = %d\n", cap(num3))

	// none
	num4 := numbers[:]
	fmt.Println("none")
	fmt.Printf("num4 = %v\n", num4)
	fmt.Printf("length = %d\n", len(num4))
	fmt.Printf("capacity = %d\n", cap(num4))
}
