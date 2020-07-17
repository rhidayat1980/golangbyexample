package main

import "fmt"

func main() {
	// if statement in go
	a := 6
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	b := 4
	if b > 3 && b < 6 {
		fmt.Println("b is within range")
	}

	// if else statement in go
	c := 1
	d := 2

	if c > d {
		fmt.Println("c is greater than d")
	} else {
		fmt.Println("d is greater than c")
	}

	// if else ladder in go
	age := 29
	if age < 18 {
		fmt.Println("Kid")
	} else if age >= 28 && age < 40 {
		fmt.Println("Young")
	} else {
		fmt.Println("Old")
	}

	// nested if else in go

	a = 1
	b = 2
	c = 3
	if a > b {
		if a > c {
			fmt.Println("Biggest is a")
		} else if b > c {
			fmt.Println("Biggest is b")
		}
	} else if b > c {
		fmt.Println("Biggest is b")
	} else {
		fmt.Println("Biggest is c")
	}

	// if with short statement in go

	if a = 6; a > 5 {
		fmt.Println("a is greater than 5")
	}

	if a = 1; a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}

	// if conditional statement
	if value := true; value == true {
		fmt.Println("value is true")
	}
}
