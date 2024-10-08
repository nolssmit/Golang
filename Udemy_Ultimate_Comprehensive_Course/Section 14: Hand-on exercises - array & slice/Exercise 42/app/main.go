package main

import (
	"fmt"
)

func main() {
	// var a[5] int
	// a[0] = 0
	// a[1] = 2
	// a[2] = 4
	// a[3] = 6
	// a[4] = 8

	a := [5]int{}

	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	for i, v := range a {
		fmt.Printf("value: %v - type: %T - index: %v\n", v, v, i)
	}
}
