package main

import (
	"fmt"
)

func myfunc2() func() {

	f := func() {
		fmt.Println("This is the returned function")
	}

	return f

}

func main() {

	t := myfunc2()
	t()

}
