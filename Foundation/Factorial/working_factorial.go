package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Request the number:
	fmt.Print("Enter a positive number between 0 and 10: ")

	// Run loop to read the number:
	for {
		var input string
		fmt.Scanln(&input)

		// Conver the input (string) from user to integer:
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error: ", err)
		}
		if len(input) == 0 {
			print("Please, enter a value")
			break
		}
		if inputInt <= 0 || inputInt > 100 {
			panic("PANIC !!!!!!!!")
		} else {
			factorial := 1
			for i := 1; i <= inputInt; i++ {
				factorial *= i
			}
			fmt.Print(factorial)
		}

	}
}
