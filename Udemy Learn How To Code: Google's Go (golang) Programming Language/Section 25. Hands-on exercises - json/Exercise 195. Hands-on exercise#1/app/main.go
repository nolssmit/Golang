package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Exercise data
// {"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"
func main() {
	// The struct fields must begin wit a capital letter because
	// they are exported in the json.Marshall function
	u1 := user{
		First:   "James",
		Last:    "Bond",
		Age:     32,
		Sayings: []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In his majesty's royal service"},
	}
	u2 := user{
		First:   "Miss",
		Last:    "Moneypenny",
		Age:     27,
		Sayings: []string{"James, it is soo good to see you", "Would you like me to take care of that for you, James?", "I would really prefer to be a secret agent myself."},
	}
	u3 := user{
		First:   "M",
		Last:    "Hmmmm",
		Age:     54,
		Sayings: []string{"Oh, James. You didn't.", "Dear God, what has James done now?", "Can someone please tell me where James Bond is?"},
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)
	fmt.Println("------------json data with three objects, each object between {}. Ech object has a name and value pair ---------------")
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
