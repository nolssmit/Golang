package main

import (
	"fmt"
)

func main() {
	m := make(map[string] [] string) // map has value of a slice of string
	m["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m["flemming_ian"] = []string{"steaks", "cigars", "espionage"}

	fmt.Println("=====================================================================")
	fmt.Printf("%#v\n", m)
	fmt.Println("=====================================================================")

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(" ")
		for i, v2 := range v {  // the value is the slice of string
			fmt.Println(i, v2)
		}
		fmt.Println("-------------------------------------------")
	}
	fmt.Println("================Delete a record from the map========================")
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(" ")
		for i, v2 := range v {  // the value is the slice of string
			fmt.Println(i, v2)
		}
		fmt.Println("-------------------------------------------")
	}
}