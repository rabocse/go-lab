package main

import (
	"fmt"
)

func foo() {
	fmt.Println("Printing a number...")
}

func foo1() int {

	return 100
}

func bar() (x int, y string) {

	x = 100
	y = "hundred"

	return x, y

}

func main() {

	foo()
	x := foo1()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)
}
