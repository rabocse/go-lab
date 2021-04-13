// Practice with multiple switch statements. Examples taken from gobyexample.com/switch

package main

import (
	"fmt"
	"time"
)

func main() {

	// This is a basic switch:

	fmt.Println(" ")
	fmt.Println("########################## BASIC SWITCH EXAMPLE ##################################")
	fmt.Println(" ")

	i := 3
	fmt.Print("Write ", i, " as ")

	switch i {

	case 1:
		fmt.Println("\"uno\"")
	case 2:
		fmt.Println("\"dos\"")
	case 3:
		fmt.Println("\"tres\"")
	}

	fmt.Println(" ")
	fmt.Println("########################### \"DEFAULT\" KW USED IN SWTICH EXAMPLE ##################")
	fmt.Println(" ")

	//  This is switch that uses "default" case.
	// You can use commas to separate multiple expressions in the same "case" statement.

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("It's WEEKEND !!!!!!!!!!!!")
	default:
		fmt.Println("It's a work day :( ")
	}

	fmt.Println(" ")
	fmt.Println("########################### SWITCH WITHOUT AN EXPRESSION ############################")
	fmt.Println(" ")

	t := time.Now()
	switch {

	case t.Hour() < 12:
		fmt.Println("It is before noon")
	case t.Hour() > 12:
		fmt.Println("It is after noon")
	}

	fmt.Println(" ")
	fmt.Println("##################### TYPE SWITCH TO COMPARE TYPES INSTEAD OF VALUES ##################")
	fmt.Println(" ")

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Dont know the type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1990)
	whatAmI("Hey")

}
