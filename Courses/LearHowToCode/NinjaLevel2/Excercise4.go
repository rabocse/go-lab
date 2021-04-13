package main

import (
	"fmt"
)

var number int

func main() {

	number = 255

	// Prints the number in decimal, binary and hexadecimal.
	fmt.Printf("%d\t %b\t %x\t", number, number, number)

	shiftedNumber := number << 2

	fmt.Println("")

	fmt.Printf("%d\t %b\t %x\t  ", shiftedNumber, shiftedNumber, shiftedNumber)

}

/*

Changing to the left or right and the amount of movements.

C:\Users\aleescob\code\go\src\github.com\go-lab\NinjaLevel2 (main -> origin)
λ go run Excercise4.go
255      11111111        ff
510      111111110       1fe
C:\Users\aleescob\code\go\src\github.com\go-lab\NinjaLevel2 (main -> origin)
λ
C:\Users\aleescob\code\go\src\github.com\go-lab\NinjaLevel2 (main -> origin)
λ
C:\Users\aleescob\code\go\src\github.com\go-lab\NinjaLevel2 (main -> origin)
λ go run Excercise4.go
255      11111111        ff
1020     1111111100      3fc


*/
