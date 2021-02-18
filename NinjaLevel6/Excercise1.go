package main

import (
	"fmt"
)

func foo() int {
	return 1990
}

func bar() (n int, s string) {

	n := 100 + 100

	return n, s
}

func main() {

	number, test := bar()
	fmt.Println(number)
	fmt.Println(test)
}
