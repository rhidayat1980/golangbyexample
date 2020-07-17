package main

import "fmt"

func main() {

	sample1 := [2]string{"a", "b"}
	fmt.Printf("Sample1 before: %v\n", sample1)

	sample2 := sample1
	sample2[0] = "c"

	fmt.Printf("Sample1 after assignment: %v\n", sample1)

	fmt.Printf("Sample2: %v\n", sample2)

	test(sample1)

	fmt.Printf("Sample1 Afer Test function call: %v\n", sample1)
}

func test(sample [2]string) {
	sample[0] = "d"
	fmt.Printf("Sample in Test function: %v\n", sample)
}
