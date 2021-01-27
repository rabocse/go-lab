package main

import "fmt"

var x, y, z int

func main() {

	// Assign different values to x, y and z to see the different results of the conditionals.

	x = 300
	y = 300
	z = 300

	if x == 200 {

		fmt.Println("They are the same !")
	}

	if x <= y {

		fmt.Println("\"y\" is greater or equal than \"x\"")

	}

	if z >= y {

		fmt.Println("\"z\" is greater or equal than \"y\"")

	}

	if x != y {

		fmt.Println("\"x\" and \"y\" are different")
	} else {

		fmt.Println("They are the same number !")
	}

	if x < y {

		fmt.Println("\"x\" is less than \"y\"")
	} else if x > y {

		fmt.Println("\"x\" is greater than \"y\"")
	} else {

		fmt.Println("\"x\" and \"y\" have the same value")
	}

	// Evaluating expressions to then assign with ":=" which will make the program to imply the type. In this case "bool" type"
	a := (2021 == 2021)
	b := (2021 < 2022)
	c := (2021 > 2022)
	d := (2022 >= 2023)

	fmt.Println(a, b, c, d)
}
