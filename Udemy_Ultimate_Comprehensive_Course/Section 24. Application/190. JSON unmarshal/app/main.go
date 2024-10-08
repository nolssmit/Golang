package main

import (
	"encoding/json"
	"fmt"
)

// Convert json data to Go
// https://mholt.github.io/json-to-go/
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// Create a string variable containing the json data
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	fmt.Println("---------------------------")

	// Unmarshall the data
	// https://pkg.go.dev/encoding/json#Unmarshal
	// https://pkg.go.dev/github.com/goccy/go-json#Unmarshal
	// https://medium.com/@homayoonalimohammadi/handling-json-in-go-804a1d9bddf5

	// people := []person{}
	var people = []person{}
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All of the data:", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)   // index
		fmt.Println(v.First, v.Last, v.Age) // value
	}
}

/*
  Slice data
  [{James Bond 35} {Miss Moneypenny 27}]
  
  JSON Data
  `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
*/
