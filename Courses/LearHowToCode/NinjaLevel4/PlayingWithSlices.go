package main

import "fmt"

var america []string
var europe = []string{"Spain", "Portugal", "Italy", "France", "Germany", "Poland"} // <-------- !!!
var continents [][]string

var test int = 3

func main() {

	fmt.Println(" ")
	fmt.Println("################################# TWO-DIMENSIONAL SLICE    ################################################## ")
	fmt.Println(" ")

	escobars := []string{"Alvaro", "Dario", "Lourdes", "Alix", "Gladis", "Eloy"}

	salazars := []string{"Marta", "Nady", "Angelica", "Pedro", "Eduardo"}

	fmt.Println(escobars)
	fmt.Println(salazars)

	people := [][]string{escobars, salazars}

	fmt.Println(people)

	fmt.Println(" ")
	fmt.Println("################################# PLAYING WITH SLICES  ################################################## ")
	fmt.Println(" ")

	asia := make([]string, 10)

	fmt.Printf("The length of \"america slice\" is: %v\n", len(america))
	fmt.Printf("The length of \"europe slice\" is: %v", len(europe))
	fmt.Println(" ")
	fmt.Printf("The capacity of \"america slice\" is: %v\n", cap(america))
	fmt.Printf("The capacity of \"europe slice\" is: %v", cap(america))
	fmt.Println(" ")
	fmt.Println("Their lenght and capacity is 0 because they were created like this: var america []string ")
	fmt.Println(" ")
	fmt.Printf("The length of \"asia slice\" is: %v\n", len(asia))
	fmt.Printf("The capacity of \"asia slice\" is: %v", cap(asia))
	fmt.Println(" ")
	fmt.Println("The lenght and capacity is 10 for both because \"asia\" was created like this: asia := make([]string, 10) ")

	fmt.Println(" ")
	america = append(america, "Canada", "US", "Mexico", "Panama", "Colombia", "Argentina", "Brasil", "Uruguay")
	fmt.Println("And now, the lenght and capacity changed because the \"append\" function was used ")
	fmt.Println(len(america))
	fmt.Println(cap(america))
	fmt.Println(america)

}
