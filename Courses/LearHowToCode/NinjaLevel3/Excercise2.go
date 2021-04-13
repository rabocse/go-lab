// Print every rune code point of the uppercase alphabet three time. It should look like this:
/*

65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
 … through the rest of the alphabet characters

*/

package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {

		fmt.Printf("%d \n", i)

		for x := 1; x <= 3; x++ {
			fmt.Printf("\t%U\t %c\n", i, i) // 			fmt.Printf("\t%#U\n", i) would do the trick as well.

		}
	}
}
