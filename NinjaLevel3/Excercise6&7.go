// Create a program that uses "if", "else if" and "else"

package main

import (
	"fmt"
)

func main() {

	Alex := 1990
	Magda := 1991

	if Magda == Alex {
		fmt.Println("Magda and Alex are the same age")
	} else if Magda > Alex {

		fmt.Println("Magda is younger than Alex")
	} else {

		fmt.Println("Magda is older than Alex")
	}

}
