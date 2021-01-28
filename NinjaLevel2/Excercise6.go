//Use iota to print 4 costants for the last for years.

package main

import (
	"fmt"
)

const (
	a = 2018 + iota
	b = a + iota
	c = a + iota
	d = a + iota
)

func main() {

	fmt.Printf("%v\t%v\t%v\t%v\t", a, b, c, d)

}
