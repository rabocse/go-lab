package main

import (
	"fmt"
)

func foo(xi ...int) int {

	total := 0
	for _, v := range xi {
		total += v
	}

	return total

}

func main() {

	x := foo(2, 3, 4, 5)
	fmt.Println(x)

}
