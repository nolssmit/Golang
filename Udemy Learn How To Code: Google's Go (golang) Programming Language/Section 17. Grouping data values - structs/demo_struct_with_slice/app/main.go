package main

import "fmt"

type favFlavor []string

type aPerson struct {
	first string
	//	flavor []string
	flavor favFlavor
}

func main() {
	p1 := aPerson{
		first:  "Nols",
		flavor: favFlavor{"chocolate", "strawberry"},
	}
	fmt.Println(p1)
}
