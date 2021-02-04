/* Using a COMPOSITE LITERAL:
- Create a slice of type int.
- Assign 10 values
- Range over the slice and print the values out.
- Print out the type of the slice.

*/

package main

import (
	"fmt"
)

func main() {

	s := []int{18,
		1,
		2015,
		29,
		4,
		1990,
		31,
		5,
		1993,
		14,
		9,
		1990}

	for i, v := range s {

		fmt.Printf("%v\t%T\t%v\n", i, v, v)
	}
}
