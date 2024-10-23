// Package mymath provides ACME inc math functions
package main	

import (
	"fmt"
)

// Sum adds an unlimited number of values of type int
func Sum(xi ...	int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}	

func main() {	
	// sum := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	sum := Sum(1, 2, 3)
	fmt.Println("Sum:", sum)
}