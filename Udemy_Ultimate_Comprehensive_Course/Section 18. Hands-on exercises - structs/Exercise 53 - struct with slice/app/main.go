package main

import (
	"fmt"
)

type favFlavor []string

type person struct {
	first string
	last  string
	flavor favFlavor 
}

func main() {
	p1 := person{
		first:  "Nols",
		last:   "Smit",
		flavor: favFlavor {"chocolate", "strawberry"},
	}

	p2 := person{
		first:  "Pat",
		last:   "Smit",
		flavor: []string{"strawberry", "butter pecan"},
	}

	fmt.Printf("%#v\n%#v\n", p1, p2)

	fmt.Println(p1)

	fmt.Printf("\n%v %v likes ", p1.first, p1.last)
	for _, v := range(p1.flavor) {
		fmt.Printf("%v, ", v)
	}
	
	fmt.Printf("\n%v %v likes ", p2.first, p2.last)
	for _, v := range(p2.flavor) {
		fmt.Printf("%v, ", v)
	}	
}
