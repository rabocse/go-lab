package main

import (
	"fmt"
)

func duplicatePrinting(f func()) (a func(), b func()) {

	var func1 func() = f
	var func2 func() = f

	return func1, func2

}

var fx func() = fmt.Println("test")

func main() {

	a, b := duplicatePrinting(fx)

}

// Wrong  ! Need to fix this.
