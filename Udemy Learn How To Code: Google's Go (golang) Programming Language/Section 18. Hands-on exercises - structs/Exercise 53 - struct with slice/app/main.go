package main

import (
	"fmt"
)

type favFlavor [] string

type person struct {
	first string
	last string
	fav favFlavor //[] string {"chocolate, "strawberry", "butter pecan", matcha"}
}



func main() {
	p1 := person {
		first: "Nols",
		last: "Smit",
		flavor: favFlavor{"chocolate, "strawberry"},
	}

	p1 := person {
		first: "Pat",
		last: "Smit",
		flavor: favFlavor{"strawberry", "butter pecan"}
	}	
}