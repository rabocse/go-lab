// Create and print a map. The key will be the name and the value its favorite thing.

package main

import "fmt"

func main() {

	m := make(map[string][]string) // MAP of string TO slice of string

	// This is how not do assign vales to the map: https://play.golang.org/p/-fudOVE4wCV

	// These are multiple ways to do it:

	// https://play.golang.org/p/RHHvMmW675R

	// https://play.golang.org/p/46vg1kr-91n

	// https://play.golang.org/p/T0GW40KWgqz

	m["Mr Deadpool"] = []string{"killing", "shooting", "stabbing"}

	m["Jean Grey"] = []string{"helping", "teaching", "burning"}

	m["Quicksilver"] = []string{"running", "stealing", "joking"}

	fmt.Printf("\tTHIS IS THE COMPLETE MAP: %v\n", m)

	fmt.Println(" ")

	for key, value := range m {

		fmt.Printf("KEY: %v\t\n", key)

		for _, s := range value {

			fmt.Printf("\t\t\tVALUE: %v\t\n", s)
		}
	}

}
