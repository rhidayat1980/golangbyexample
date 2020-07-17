package main

import "fmt"

func main() {

	sample := []int{}
	fmt.Println(len(sample))
	fmt.Println(cap(sample))
	fmt.Println(sample)

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)

	numbers := [5]int{1, 2, 3, 4, 5}

	// slicing start:end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1: %d\n", num1)
	fmt.Printf("length: %d\n", len(num1))
	fmt.Printf("capacity: %d\n", cap(num1))

	// only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num2: %d\n", num2)
	fmt.Printf("length: %d\n", len(num2))
	fmt.Printf("capacity: %d\n", cap(num2))

	// only end
	num3 := numbers[:4]
	fmt.Println("\nOnly end")
	fmt.Printf("num3: %d\n", num3)
	fmt.Printf("length: %d\n", len(num3))
	fmt.Printf("capacity: %d\n", cap(num3))

	// none
	num4 := numbers[:]
	fmt.Println("\nNone")
	fmt.Printf("num4: %d\n", num4)
	fmt.Printf("length: %d\n", len(num4))
	fmt.Printf("capacity: %d\n", cap(num4))

	// create a slice with make()

	numbero := make([]int, 3, 4)
	fmt.Printf("numbers: %d\n", numbero)
	fmt.Printf("length: %d\n", len(numbero))
	fmt.Printf("capacity: %d\n", cap(numbero))

	// using the new function

	numberox := new([]int)
	fmt.Printf("numberox: %d\n", *numberox)
	fmt.Printf("length: %d\n", len(*numberox))
	fmt.Printf("capacity: %d\n", len(*numberox))

}
