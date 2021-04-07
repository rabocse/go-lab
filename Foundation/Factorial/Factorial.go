package main

import (
	"fmt"
	"strconv"
)

func myfactorial(inputInt int) int {

	if inputInt <= 0 {
		panic("PANIC #1 ! Factorial cannot be calculated for negative values or zero. Please, enter a positive number")
	} else if inputInt > 100 {
		panic("PANIC #2 ! Highest factorial to calculate is 100. Please, do not enter values higher than 100.")
	} else {
		factorial := 1
		for i := 1; i <= inputInt; i++ {
			factorial *= i
		}
		return factorial
	}
}

func main() {

	// Anouynimous function to be deferred and be used later with recover()
	defer func() {
		//		fmt.Println("Deferring the function to then be used with \"recover\"")
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic scenario:", r)
		}
	}()

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
