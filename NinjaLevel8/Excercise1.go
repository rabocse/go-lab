package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {

	u1 := user{
		First: "Homer",
		Age:   40,
	}

	u2 := user{
		First: "Bart",
		Age:   10,
	}

	u3 := user{
		First: "Lisa",
		Age:   8,
	}

	users := []user{u1, u2, u3}

	fmt.Println("###################### Go Data Structure ####################################################")
	fmt.Println(users)
	fmt.Println(" ")

	fmt.Println("###################### JSON (Marshaled) ######################################################")

	m, err := json.Marshal(users) // Keep in mind that Marshal returns a []byte, so those will need to be converted to string to them print them.
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(m)) // Converting []byte to string to then print.
}
