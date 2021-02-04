// Create an ARRAY which holds 5 values of TYPE int.
// Assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing, print out the TYPE of the array.

package main

import (
	"fmt"
)

var y [10]int

func main() {

	// y := [10]int{}       // If I declare this "y variable" and I do not use it, then the compilation throws and error, as expected. But, if the variable is declared globaly, then the error is not seen. For example: var y [10]int
	x := [5]int{}

	fmt.Println("")
	fmt.Println("########################## PRINTING ZERO VALUE FOR THE ARRAY X ################################## ")
	fmt.Println("")
	fmt.Println(x)

	fmt.Println("")
	fmt.Println("####### ASSIGNING VALUES AND PRINTING  EACH VALUE FOR THE INDEX POSITIONS OF THE ARRAY X ############### ")
	fmt.Println("")

	x[0] = 18
	x[1] = 01
	x[2] = 2015
	x[3] = 4
	x[4] = 29

	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	fmt.Println("")
	fmt.Println("########################## RANGING OVER THE ARRAY ################################## ")
	fmt.Println("")

	for _, v := range x {

		fmt.Printf("%T\t %v\n", v, v)

	}
}
