package main

import (
	"fmt"
)

// Define the interface with two methods:
type iGeoshape interface {
	Draw()
	Print()
}

// Define Rectangle type:
type Rectangle struct {
}

// Define methods for Rectangle:
func (Rectangle) Draw() {
	fmt.Println("Drawing a rectagle...")
}

func (Rectangle) Print() {
	fmt.Println("Printing a rectagle...")
}

// Define Circle type:
type Circle struct {
}

// Define methods for Circle:
func (Circle) Draw() {
	fmt.Println("Drawing a circle...")
}

func (Circle) Print() {
	fmt.Println("Printing a circle...")
}

// Render function:
func Render(a ...iGeoshape) {
	for _, v := range a {
		v.Draw()
	}
}

func main() {

	var rect1 Rectangle
	var rect2 Rectangle
	var rect3 Rectangle
	var circ1 Circle
	var circ2 Circle

	// Render accepts Rectangle and Circle
	Render(rect1, rect2, rect3, circ1, circ2)

}
