// Pring a "raw string"

package main

import (
	"fmt"
)

// This is  a raw string
var rawString string = ` This is a raw string. 

To be defined the "backtick" needs to used instad of double quotes ("") `

func main() {

	fmt.Println(rawString)

}
