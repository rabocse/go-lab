package main

import (
	"fmt"
	"os"
)

func main() {

	// Create a map with key and values for airport codes

	airport := map[string]string{"Atlanta": "ATL",
		"Beijing": "PEK",
		"Chicago": "ORD"}

	fmt.Println(airport) // Print to verify the map.
	fmt.Println("WE ARE GOOD SO FAR")

	// Accept and airport codes (KEYS) as command arguments.

	code1 := os.Args[1]
	code2 := os.Args[2]
	code3 := os.Args[2]

	fmt.Printf("%v \n", code1)
	fmt.Printf("%v \n", code2)
	fmt.Printf("%v \n", code3)

	// Validate three codes are entered

	if len(os.Args) == 4 {
		// fmt.Println("Ok ok")
	} else {
		fmt.Println("Entry full intinerary (3 codes)")
	}

	// Compare entered code against map:


	for k, _ := range airport {


		_, exist = airport
		k, ok =airport[key]

		fmt.Println(k)


	

}
