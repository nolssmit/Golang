// https://www.udemy.com/course/go-language/learn/lecture/35216188#notes
//
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- FOR LOOPS -----
	// for initialization; condition; postStatement {BODY}
	// Print numbers 1 through 5
	pl("-----------increment-----------")
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Not in scope outside loop therefore an error
	// pl(x)

	// Do the opposite
	pl("-----------decrement-----------")
	for x := 5; x >= 1; x-- {
		pl(x)
	}


	// x is out of the scope of the for loop so it doesn't exist
	// pl("x :", x)

	// For is used to simulate while loops as well	
	pl("--------Simulate a while loop--------")
	fx := 0
	for fx < 5 {
		pl(fx)
		fx++
	}

	// Cycle through an array with range
	// More on arrays later
	// We don't need the index so we ignore it
	// with the blank identifier _
	pl("---------Cycle thru an array of numbers-------")
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

	// We can allow a condition in the for loop
	// to decide when to exit
	pl("------Conditionally cycle thru an array of numbers-----")
	xVal := 1
	for true {
		if xVal == 5 {
			break
		}
		pl(xVal)
		xVal++
	}

}
