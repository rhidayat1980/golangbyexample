package main

import "fmt"

func main() {

	// Both number of elements and actual elements
	sample1 := [2]int{1, 2}
	fmt.Printf("Sample1: %v Len: %d\n", sample1, len(sample1))

	// only actual elements
	sample2 := [...]int{1, 2}
	fmt.Printf("Sample2: %v Len: %d\n", sample2, len(sample2))

	// only number of elements
	sample3 := [2]int{}
	fmt.Printf("Sample3: %v Len: %d\n", sample3, len(sample3))

	// without both number of elements and actual elements
	sample4 := [...]int{}
	fmt.Printf("Sample4: %v Len: %d\n", sample4, len(sample4))

	sample5 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("Sample5: %v Len: %d\n", sample5, len(sample5))

	sample6 := [4]int{5, 8}
	fmt.Printf("Sample6: %v Len: %d\n", sample6, len(sample6))
}
