// Print a number in binary, decimal and hex.package NinjaLevel2

package main

import (
	"fmt"
)

var x int

func main() {

	x = 255

	fmt.Printf("%b \n", x)
	fmt.Printf("%x \n", x)
	fmt.Printf("%X \n", x)
	fmt.Printf("%d \n", x)
	fmt.Printf("%d\t%b\t%#x", x, x, x)

}
