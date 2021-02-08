package main

import (
	"fmt"
)

// Slice declared and assigned globally.
var europe = []string{"Spain", "Portugal", "Germany", "Poland", "Italy", "France"}

var america = []string{"Canada", "US", "Mexico", "Puerto Rico", "Costa Rica", "Panama", "Colombia", "Argentina", "Brasil", "Chile"}

func main() {

	// Slice created with "short-declaration"
	merge := [][]string{america, europe}

	fmt.Printf("This is america: %v\n", america)
	fmt.Printf("This is europe: %v\n", europe)
	fmt.Printf("This is a merge of america and europe: %v\n", merge)

	fmt.Println("############################### #########################################\n")

	for i, member := range merge { // Goes through merge which contains two slices. Then, it gets the index (i) for each slice AND the slice itself (member)

		fmt.Printf("record %v: \n", i) // Then, we are printing only the index (i)...But, remember, we have now "member" that is taking the value of each slice as it loops.
		fmt.Println(" ")
		for j, memberSlice := range member { // So, we loop over "member" and we get the index (we call it now j) and the values from the slice.

			fmt.Printf("\tIndex position %v \t  Value: %v \n", j, memberSlice) // Finally, we print both, the index (j) and the slice itself (memberSlice)

		}
	}
}
