package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("########################### JSON (Before unmarshaling) #######################################################")
	fmt.Println(" ")
	fmt.Println(s)
	fmt.Println(" ")
	fmt.Println(" ")

	var m interface{}

	bs := []byte(s)

	err := json.Unmarshal(bs, &m)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("########################### Go Data Structure  (After unmarshaling) #######################################################")
	fmt.Println(" ")
	fmt.Println(m)

}
