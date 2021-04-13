package main

import (
	"fmt"
)

func main() {

	m := make(map[string][]string)

	m["Mr Deadpool"] = []string{"killing", "shooting", "stabbing"}

	m["Jean Grey"] = []string{"helping", "teaching", "burning"}

	m["Quicksilver"] = []string{"running", "stealing", "joking"}

	// Printing the map with "original" values:
	fmt.Println("###### INITIAL MAP BEFORE ADDING ANOTHER KEY:VALUE #####\n")
	fmt.Println(m)
	fmt.Println(" ")

	// Adding a new record for my map:
	m["Magneto"] = []string{"flying", "metal-bending", "planning"}

	// Printing the lengh of my map:
	fmt.Println("################### MAP LENGTH ########################\n")
	fmt.Println(len(m))
	fmt.Println(" ")

	// Printing the map after adding a record:
	fmt.Println("####### MAP AFTER ADDING \"Magneto\" Key ##############\n")
	fmt.Printf("THIS IS THE  WHOLE MAP: %v\n", m)
	fmt.Println(" ")

	// Looping over the map printing the key and value:
	fmt.Println("################### LOOPING OVER MAP #####################\n")
	for key, value := range m {

		fmt.Printf("%v:%v\n", key, value)
	}
	fmt.Println(" ")

	// Deleting a record  and then printing the map:
	fmt.Println("################### \"Quicksilver\" KEY DELETED #####################\n")
	delete(m, "Quicksilver")
	for key, value := range m {

		fmt.Printf("%v:%v\t", key, value)
	}

}
