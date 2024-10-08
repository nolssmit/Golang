package main

import (
	"fmt"
)

/*
Here's how you check if a map contains a key.

val, ok := myMap["foo"]
// If the key exists
if ok {
    // Do something
}
This initializes two variables. val is the value of "foo" from the map if it exists, or a "zero value" if it doesn't (in this case the empty string). ok is a bool that will be set to true if the key existed.

If you want, you can shorten this to a one-liner.

if val, ok := myMap["foo"]; ok {
    //do something here
}
*/

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key: %v, \t value: %v\n", k, v)
	}

	fmt.Println("----------------------")

	age1 := m["James"]
	fmt.Println("The age of Bond:", age1)

	fmt.Println("----------------------")

	if v, ok := m["James"]; ok {
		fmt.Println("There is a Bond lookup entry, and Bond's age:", v)
	}

	fmt.Println("----------------------")

	age2 := m["Q"]
	fmt.Println("The age of Q:", age2)

	fmt.Println("----------------------")

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is zero value of int:", v)
	}
}
