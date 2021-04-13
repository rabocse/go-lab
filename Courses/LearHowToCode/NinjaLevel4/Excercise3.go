package main

import (
	"fmt"
)

func main() {

	// Create and assign values to an slice s with a composite literal.
	s := []int{4, 29, 1990, 14, 9, 1990, 18, 01, 2015}

	// Loop over the slice s with "range" keyword
	for i, v := range s {
		fmt.Printf("%v\t%T\t%v\n", i, v, v)
	}

	// Loop over the slice "s" with regular "for" (no range keyword)
	for i := 0; i <= len(s); i++ {

		fmt.Println(i, s)
	}

	fmt.Println(" ")
	fmt.Println(" ###################### USING SLICING .... ########################################")
	fmt.Println(" ")
	/* Using SLICING to create the following new slices which are then printed

	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]


	*/

	biggerSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} // First, I created a new slice called "biggerSlice"

	s = append(s, biggerSlice...) // Then, biggerSlice is appended to the slice called "s". If the ... are not used, then the append function might interpret biggerSlice as int and produce an error when compiling.

	fmt.Println(s)

	// This is SLICING !!!!!!!!!!!!!!!!!!
	fmt.Println(s[9:14])
	fmt.Println(s[15:20])
	fmt.Println(s[11:16])
	fmt.Println(s[10:15])

}
