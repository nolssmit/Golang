package main

import (
	"encoding/json"
	"fmt"
)

// Use the online app to get the format of the struct: https://mholt.github.io/json-to-go/
type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	// json data
	jsondata := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	bs := []byte(jsondata)

	fmt.Printf("json data of type: %T\n", jsondata)
	fmt.Println(jsondata)
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Slice data of type: %T\n", bs)

	// Unmarshall the data
	// https://pkg.go.dev/encoding/json#Unmarshal
	// https://pkg.go.dev/github.com/goccy/go-json#Unmarshal
	// https://medium.com/@homayoonalimohammadi/handling-json-in-go-804a1d9bddf5

	var users = []user{}
	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
	fmt.Println("-----------------------------------------------------------")
	for i, user := range users {
		fmt.Println("Person #", i)
		fmt.Println("\t", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t,\t,\t", saying)
		}
	}
}
