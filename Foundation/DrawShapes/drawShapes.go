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

// func render( a []iGeoshape) {

func main() {

}
