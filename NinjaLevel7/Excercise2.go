package main

import "fmt"

type person struct {
	name     string
	lastname string
}

// See how the current code and the commented one work both  to accomplish the mutation of the data. HOWEVER, notice that the return was indeed not needed !

func changeMe(p *person) person {
	// func changeMe(p *person) *person {

	p.name = "Pedro"
	p.lastname = "Pablo"
	// return p
	return *p
}

func main() {

	p1 := person{
		name:     "Pawel",
		lastname: "Piotr",
	}

	fmt.Println("################# Before Using changeMe() #########################################")
	fmt.Println("   ")
	fmt.Println(p1.name)
	fmt.Println(p1.lastname)
	fmt.Println("   ")
	fmt.Println("################# After Using changeMe() #########################################")
	fmt.Println("   ")
	changeMe(&p1) // I did no take and assign the return value from changeMe, I just called the fuction !
	fmt.Println(p1.name)
	fmt.Println(p1.lastname)

}
