package main

import "fmt"

type person struct {
	firstname    string
	lastname     string
	favIceCreams []string
}

func main() {

	p1 := person{

		firstname:    "Ronaldo",
		lastname:     "Nazario",
		favIceCreams: []string{"vanilla", "banana", "chocolate"},
	}

	p2 := person{
		firstname:    "Lionel",
		lastname:     "Messi",
		favIceCreams: []string{"Chicle", "limon", "mate"},
	}

	m := map[string]person{p1.lastname: p1, p2.lastname: p2}

	fmt.Println(m)

	fmt.Println(" ")
	fmt.Println("##########################################################################################################")
	fmt.Println(" ")

}
