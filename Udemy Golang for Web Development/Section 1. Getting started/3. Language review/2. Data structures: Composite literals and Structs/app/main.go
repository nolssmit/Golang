package main

import (
	"fmt"
)

// Declaring your own type
type hotdog int

// Declaring a struct
// Struct is a collection of fields
type Person struct {
	// Lowercase field names are private, not visible outside package
	fname string
	lname string
}

func main() {
	var t hotdog
	t = 7
	fmt.Printf("Type of t is: %T\n", t)
	fmt.Printf("Value of t is: %v\n\n", t)

	// Slice of integers
	// Slice is a data structure used to hold a list of values
	// You don't need a trailing comma when you list it out in one row
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println("Value of slice xi is:", xi)

	// You need a trailing comma when you list it out in mote than one row
	// Use maps for key-value pairs and quick lookups
	m := map[string]int{
		"foo": 1,
		"bar": 2,}
	fmt.Println("Value of map m is:", m)

	// Structs. You don't need to specify field names
	// when you assign values.
	p1 := Person{"John", "Doe"}
	fmt.Println("Value of struct p1 is:", p1)
}
