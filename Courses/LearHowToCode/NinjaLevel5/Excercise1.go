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

	fmt.Printf("\t\tThe name of the first person is: %v\n", p1.firstname)
	fmt.Printf("\t\tThe lastname of the first person is: %v\n", p1.lastname)

	for _, v := range p1.favIceCreams {

		fmt.Printf("\t\tThese is one of %v's favourite icecreams: %v\n", p1.firstname, v)
	}

	fmt.Println(" ")
	fmt.Println("##########################################################################################################")
	fmt.Println(" ")

	p2 := person{
		firstname:    "Lionel",
		lastname:     "Messi",
		favIceCreams: []string{"Chicle", "limon", "mate"},
	}

	fmt.Printf("\t\tThe name of the second person is: %v\n", p2.firstname)
	fmt.Printf("\t\tThe lastname of the second person is: %v\n", p2.lastname)

	for _, v := range p2.favIceCreams {

		fmt.Printf("\t\tThese is one of %v's favourite icecreams: %v\n", p2.firstname, v)
	}

	fmt.Println(" ")
	fmt.Println("##########################################################################################################")
	fmt.Println(" ")

}
