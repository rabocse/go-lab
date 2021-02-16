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

	m := map[string]person{p1.lastname: p1,
		p2.lastname: p2}

	for k, v := range m {
		fmt.Printf("\t%v:\n", k)

		fmt.Printf("\t\t\t%v\n", v.firstname)
		fmt.Printf("\t\t\t%v\n", v.lastname)

		for i, val := range v.favIceCreams { // dot-notation can be used since v is person type.

			fmt.Println(i, val)
		}

	}

}
