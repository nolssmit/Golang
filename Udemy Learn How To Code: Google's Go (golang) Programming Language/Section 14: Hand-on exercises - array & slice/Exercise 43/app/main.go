package main

import (
	"fmt"
)

func main() {
		xi := [] int {42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// var xi = [5]int{}
	// for i := 0; i < len(xi); i++ {
	// 	xi[i] = 2 * i
	// }

	for i, v := range xi {
		fmt.Printf("index: %v - value: %v  - type: %T\n", i, v, v)
	}

	fmt.Printf("index: %T - value: %#v\n", xi, xi)
	fmt.Printf("index: %T - value: %v\n", xi, xi)
}
