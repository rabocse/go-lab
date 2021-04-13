package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	// Print formats using the default formats for its operands and writes to standard output.
	//Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error

	// First three "Print" statements use the default formats. Also, no "escape sequence" was used:

	fmt.Print("\n ################ Print (Default Format and Without Using \"Escape Sequences\" ) ############################# \n")
	fmt.Print("\n")

	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)

	fmt.Print("\n") // Notice how unpractical this can be when needing multiple space lines. Next example uses a more handy way.

	fmt.Print("\n ################ Print With \"Escape Sequences\" #############################") // Notice that I am not using \n here.

	// Instead of \n, I am using this "BlankLines" which is more handly when it comes to print multiple line strings. In this case is just blank.
	BlankLines := `
	



	`
	fmt.Print(BlankLines, "\n")

	/*
		 Escape Sequences:

		\n    A new line character.
		\t    A tab character.
		\"    Double quotation marks.
		\\    A backslash

	*/

	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
	fmt.Print(z, "\n")

	fmt.Print(BlankLines, "\n")

	// Printf formats according to a format specifier and writes to standard output.
	//It returns the number of bytes written and any write error encountered.
	// For example, "%v" is	the value in a default format.

	fmt.Print("\n ################ Printf =! Print  ############################# \n")
	fmt.Print(BlankLines)
	fmt.Print("\n The variables X, y and z are printed with the specified format. In this case \"%T\" which means to print the Type and not the value. \n")
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)
	fmt.Print("\n We have of course multiple \"verbs\" to get different format outputs, \"%T\" is just one of the many. \n")

	fmt.Print(BlankLines, "\n")

	//Println formats using the default formats for its operands and writes to standard output.
	//Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.

	fmt.Println("############## Println = Print + NewLine #############################")
	fmt.Println(BlankLines)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(BlankLines)
	fmt.Println(x, y, z)
	fmt.Println(BlankLines)
	fmt.Println(x, y, z)

}
