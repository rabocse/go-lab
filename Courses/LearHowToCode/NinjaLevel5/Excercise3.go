package main

import "fmt"

func main() {

	type vehicle struct {
		doors string
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: "conventional",
			color: "red"},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: "butterfly",
			color: "black"},
		luxury: true,
	}

	fmt.Println("####################################################################################")
	fmt.Println()
	fmt.Printf("TRUCK: %v\n", t)
	fmt.Println()
	fmt.Printf("SEDAN: %v\n", s)
	fmt.Println()
	fmt.Println(t.fourWheel)
	fmt.Println(s.luxury)
	fmt.Println("####################################################################################")

}
