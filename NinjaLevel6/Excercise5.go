package main

import (
	"fmt"
	"math"
)

type square struct { //Type "square" with "length" and "width" as fields.
	length float64
	width  float64
}

type circle struct { //Type "circle" with "radius" as field.
	radius float64
}

func (c circle) area() float64 { //Method called "area" to calculate the area of a variable of type circle.
	area := (math.Pi * c.radius * c.radius)

	return area
}

func (s square) area() float64 { //Method called "area" to calculate the area of a variable of type square.
	area := s.length * s.width

	return area
}

type shape interface {
	area() float64
}

func info(sh shape) {

	area := sh.area()
	fmt.Println(area)

}

func main() {

	radiusCircle := circle{ // radiusCircle is type circle.
		radius: 10.5,
	}

	areaCircle := radiusCircle.area() // areaCircle is calculated by passing radioCircle to the method AREA.
	fmt.Println(areaCircle)

	dimensions := square{ // dimensions is type SQUARE.
		length: 10.5,
		width:  100.75,
	}

	areaSquare := dimensions.area() // areaSquare is calculated by passing areaSquare to the method AREA.
	fmt.Println(areaSquare)

	info(radiusCircle)
	info(dimensions)

}
