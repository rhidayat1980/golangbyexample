package main

import "fmt"

func main() {
	printType("test_string")
	printType(10)
}

func printType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type: string")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Printf("Unknown type %T", v)
	}
}
