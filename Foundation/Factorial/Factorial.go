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

	// Request the number:
	fmt.Print("Enter a positive number: ")

	// Run loop to read the number:
	for {
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

		// Now that inputInt is used to define whether is valid value to calculate factorial:
		//		if inputInt <= 0 || inputInt > 100 {
		//			panic("PANIC !!!!!!!!")

		fact := myfactorial(inputInt)
		fmt.Print(fact)

		// fmt.Print(factorial)
	}
}
