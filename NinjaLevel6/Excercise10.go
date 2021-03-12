package main

import (
	"fmt"
)

func main() {

	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := foo() // This is also foo() but it was called by another variable called "g". So the increments done by "f" do not apply.
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int {
	x := 0 // "x" is declared outside the anonymous function where it is used.  This is a "closure".
	// Notice how the fact of being declared outside the body of the anonymous function is going to allow it to mantain the value of "x" whenever (not "whenever". See difference between "f" and "g") is incremented.
	return func() int {
		y := 10 // Insted, "y" cannot be mantained as this is incremented each time the funcion is called.
		y++
		fmt.Println(y)
		x++
		return x
	}
}
