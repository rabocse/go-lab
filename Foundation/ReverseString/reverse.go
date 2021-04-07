package main

import "fmt"

func reverseText(str string) (done string) {
	for _, v := range str {
		done = string(v) + done
	}
	return
}

func main() {

	str := "Go has superior string manipulation"

	strRev := reverseText(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
