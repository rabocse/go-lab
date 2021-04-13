// Create a typed and untyped constant. Print them.package NinjaLevel2

package main

import "fmt"

const age = 30
const weight = 95.5

const (
	height   int     = 185
	siblings float64 = 2
)

func main() {

	fmt.Printf("%T is the type that the Go compiler assgined for the UNTYPED constant called \"age\" with a value of %v \n", age, age)
	fmt.Printf("%T is the type that the Go compiler assigned for the UNTYPED constant called \"weight\" with a value of %v \n", weight, weight)

	fmt.Printf("These are the statically defined types for the constants \"height\" \"siblings\": %T\t%T  ", height, siblings)

}
