package main

import "fmt"

func foo() {

	defer fmt.Println("First in code from top to bottom  but defered")

	fmt.Println("Second in code from top to bottom.")

}

func main() {

	foo()

}
