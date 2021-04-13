package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {

	f := p.first
	l := p.last
	a := p.age

	fmt.Printf("The full name person is %v %v and his/her age is %v ", f, l, a)
}

func main() {

	someone := person{
		first: "Alex",
		last:  "Escobar",
		age:   30,
	}

	someone.speak()

}
