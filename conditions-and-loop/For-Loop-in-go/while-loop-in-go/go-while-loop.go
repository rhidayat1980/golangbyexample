package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 0
	name := "rachmat"
	for i <= 10 {
		fmt.Println(strconv.Itoa(i) + "." + " " + name)
		i++
	}

	b := 1
	for b < 10 {
		fmt.Println(b)
		b++
	}
}
