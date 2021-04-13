package main

import "fmt"

type MyOwnType int

var x MyOwnType
var y int

func main() {

	//To print zero value
	fmt.Println(x)

	// To print the type of the variable
	fmt.Printf("%T \n", x)

	//Assign a value to the variable x
	x = 100

	// Print the value of x
	fmt.Println(x)

	// Use CONVERSION  to convert the TYPE of the value stored in X to the underlying type

	y = int(x)

	// Print the type of the variable y. It should be "int" since we converted x type to "int".
	fmt.Printf("%T\n", y)

}
