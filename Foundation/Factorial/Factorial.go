package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Request the number:
	fmt.Print("Enter a positive number: ")

	// Run loop to read the number:
	for {
		var input string
		fmt.Scanln(&input)

		// Convert the input (string) from user to integer:
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error: ", err)
		}
		// Validate that the user entered a value
		if len(input) == 0 {
			print("Please, enter a value")
			break
		}
		if inputInt <= 0 || inputInt > 100 {
			panic("PANIC !!!!!!!!")
		} else {
			myfactorial(inputInt)
			}
			fmt.Print(factorial)
		}

	}
}



func myfactorial(inputInt int) uint64 {  

	if inputInt <= 0 || inputInt > 100 {
		panic("PANIC !!!!!!!!")
	} else {
		factorial := 1
		for i := 1; i <= inputInt; i++ {
			factorial *= i
		}
		return factorial
	}
}
	