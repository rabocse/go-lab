package main

import (
	"fmt"
	"strconv"
)

func myfactorial(inputInt int) int {

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

func main() {

	// Run loop to read the number:
	for {
		// Request the number:
		fmt.Print("Enter a positive number: ")

		var input string
		fmt.Scanln(&input)

		// Validate that the user entered a value:
		if len(input) == 0 {
			print("Value was not entered. Exiting program...")
			break
		}
		// Convert the input (string) from user to integer and handle the error:
		inputInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Converting from string to int failed:", err)
		}
		// Calculate and print the factorial:
		fact := myfactorial(inputInt)
		fmt.Println(fact)
		fmt.Println(" ")

	}
}
