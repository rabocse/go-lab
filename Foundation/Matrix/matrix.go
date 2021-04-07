package main

import "fmt"

// mathOps is a type that follows the function pattern defined:
type mathOps func(a int, b int) int

func main() {

	// Create a slice of mathOps type:
	matrix := []mathOps{func(a int, b int) int {
		return a + b
	},
		func(a int, b int) int {
			return a / b
		},
		func(a int, b int) int {
			return a * b
		},
		func(a int, b int) int {
			return a - b
		},
	}

	// Declare new variables for each index from mathOps slice
	add := matrix[0]
	addResult := add(4, 5) // Pass parameters to each function
	fmt.Println(addResult)

	div := matrix[1]
	divResult := div(0, 3)
	fmt.Println(divResult)

	mult := matrix[2]
	multResult := mult(4, 8)
	fmt.Println(multResult)

	sub := matrix[3]
	subResult := sub(1, 4)
	fmt.Println(subResult)

}
