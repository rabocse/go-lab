package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Harcoded number to guess
	NumberToGuess := 77

	// Asking to user to enter the guess
	fmt.Print("Enter number to guess ( < 100 ): ")

	// Read the guess with infinite loop
	for {
		var guess string
		// fmt.Print("Enter text: ")
		fmt.Scanln(&guess)
		fmt.Println(guess) // Printing guess as string
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(guessInt) // Printing guess as Int
		if len(guess) == 0 {
			print("done")
			break
		} else if guessInt < NumberToGuess {
			fmt.Println("The guess was LOW")
		} else if guessInt > NumberToGuess {
			fmt.Println("The guess was HIGH")
		} else {
			fmt.Println("Your guessed it !")
			break
		}

	}
}
