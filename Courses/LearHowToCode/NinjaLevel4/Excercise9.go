package main

import (
	"fmt"
)

func main() {

	m := make(map[string][]string)

	m["Mr Deadpool"] = []string{"killing", "shooting", "stabbing"}

	m["Jean Grey"] = []string{"helping", "teaching", "burning"}

	m["Quicksilve"] = []string{"running", "stealing", "joking"}

	// Printing the map with "original" values:
	fmt.Println("###### INITIAL MAP BEFORE ADDING ANOTHER KEY:VALUE #####\n")
	fmt.Printf("\t THIS IS THE  MAP WITH THREE VALUES: %v\n", m)

	fmt.Println(" ")

	// Adding a new record for my map:
	m["Magneto"] = []string{"flying", "metal-bending", "planning"}

	// Printing the lengh of my map:
	fmt.Println("################### MAP LENGTH ########################\n")
	fmt.Println(len(m))

	// Printing the map after adding a record:
	fmt.Println("####### MAP AFTER ADDING \"Magneto\" Key ##############\n")
	fmt.Printf("\t THIS IS THE  WHOLE MAP: %v\n", m)

	// Looping over the map printing the key and value:
	fmt.Println("################### LOOPING OVER MAP #####################\n")
	for key, value := range m {

		fmt.Printf("\t%v:%v\t", key, value)
	}
}
