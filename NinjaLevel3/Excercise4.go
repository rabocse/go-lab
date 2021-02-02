// Create a loop using this syntax: for {}
// Have it preint out the years you have been alive.

package main

import (
	"fmt"
)

func main() {

	year := 1990

	for {

		if year <= 2021 { // It could be solved too by using "year > 2021" then "break". And, outside the conditional, then it is when we print and increment "year".
			fmt.Println(year)
			year++
		} else {
			break
		}
	}

}
