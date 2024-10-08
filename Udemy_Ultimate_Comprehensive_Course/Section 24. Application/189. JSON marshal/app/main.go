package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

// Note: The field names must start with a uppercase letter
// or else your output will be: [{},{}]

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	fmt.Println(p1, p2)
	fmt.Println("-------------------")
	people := []person{p1, p2}
	fmt.Println(people)
	fmt.Println("-------------------")
	// Marshall the data
	// https://pkg.go.dev/encoding/json#Marshal
	// https://pkg.go.dev/github.com/goccy/go-json#Marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

// Data for this exercise
// [{James Bond 35} {Miss Moneypenny 27}]
