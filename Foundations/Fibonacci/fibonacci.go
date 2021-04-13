package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	a := os.Args[1]
	b := os.Args[2]

	// Printing passed arguments from CLI.

	fmt.Println(a)
	fmt.Println(b)

	// I guess they are strings, so I have to convert to int.

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	starting, _ := strconv.Atoi(a)
	ending, _ := strconv.Atoi(b)

	fmt.Printf("%T \n", starting)
	fmt.Printf("%T \n", ending)

}
