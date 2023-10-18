package main

import (
	"fmt"
)

func main() {
	x := 83 / 40
	y := 83 % 40
	fmt.Printf("Modules = %v\nRemainder = %v\n", x, y)

	// Finding even numbers
	for i:=0; i<100; i++ {
		if i%2 == 0 {
			fmt.Printf("This is an even number \t%v\n", i)
		}
	}
}
