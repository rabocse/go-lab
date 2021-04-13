package main

import (
	"fmt"
	"reflect"
)

type MyOwnType int

var x MyOwnType

func main() {

	// Print out the value of varaible x
	fmt.Println(x)

	// Print out the type of the variable x
	// First I use the "reflect" package for such.
	fmt.Println(reflect.TypeOf(x))

	// Now, I use the "Printf" with the verb "%T" to get the type of the variable.
	fmt.Printf("%T \n", x)

	// Assign 42 to the variable x
	x = 42

	// Print out the valye fo the variable x
	fmt.Println(x)

}
