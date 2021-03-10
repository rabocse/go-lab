package main

import (
	"fmt"
)

// "foo" accepts two arguments:
// 1 - a funcion like "f" like func(xi []int) int
// 2 - a slice of int called "ii"
// finally, "foo" returns an int.

func foo(f func(xi []int) int, ii []int) int {

	n := f(ii) // Notice that I cannot predict the exact result from here since foo will receive a function "f" that can do whatever thing.
	// What we can expect from here is that whatever int the return from "f" is, "foo" will increment this and return it.
	n++
	return n
}

func main() {

	g := func(xi []int) int {

		if len(xi) == 0 {

			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1] // Given an slice "xi", "g" function will return the sum of the first and last value.

	}

	x := foo(g, []int{1, 2, 3, 4, 5}) // "g" and a slice are passed to "foo".
	// then, "foo" calls the same "g" and passes the slice to it.
	// then "g" returns the sum of 1 and 5 as those are the first and last values from the passed slice.
	// finally, "foo" takes (1 + 5) as return from "g" and increments 1 to it to then return it (7).
	fmt.Println(x)

}

// This is what is called a callback. A function that calls a function as parameter.
// Functions are "first class citizens" in Go.
