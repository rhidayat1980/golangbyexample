package main

import "fmt"

func main() {
	switch char := "b"; char {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
		break
		fmt.Println("after b") // unreachable code
	default:
		fmt.Println("No matching character")
	}
}
