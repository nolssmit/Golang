package main

import (
	"fmt"
)

// func main() {
// 	f := factorial(4)
// 	fmt.Println("Total: ",f)
// }

// func factorial(n int) int {
// 	total := 1
// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}
// 	return total
// }

/*
Challenge #1:
-- Use goroutines and channels to calculate factorial

Challenge #2:
-- Why might you use goroutines and channels to calculate factorial?
---- Formulate your own answer, the post that answer to this discussion area: //https://goo.gl/flGsyX
---- Read a few other answers at the discussion area to see the reasons of others
*/

func main() {
	c := factorial(4)

	// Block exiting main until channel is closed
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
// Use goroutines and channels to calculate factorial
// because it allows you to run multiple functions at the same time