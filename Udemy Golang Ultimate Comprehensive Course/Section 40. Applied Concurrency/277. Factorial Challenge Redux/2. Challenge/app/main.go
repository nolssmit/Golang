package main

import (
	"fmt"
)

// func main() {
// 	c := factorial(4)

// 	for n := range c {
// 		fmt.Println(n)
// 	}
// }

// func factorial(n int) chan int {
// 	out := make(chan int)

// 	go func() {
// 		total := 1
// 		for i := n; i > 0; i-- {
// 			total *= i
// 		}
// 		out <- total
// 		close(out)
// 	}()

// 	return out
// }

/* CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrenly and in parallel.
-- Use the "pipeline" pattern to accomplish this.

POST to DISCUSSION:
-- What realizations did you have about working with concurrent code  when building your sollution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other people's answers, here: https://goo.gl/uJa99G
*/

func main() {
	c := make(chan int)

	go factorial(c, 5)

	for n := range c {
		fmt.Println(n)
	}
}

func factorial(c chan int, n int) chan int {
	out := make(chan int)
	go func(n int, out chan int) chan int {
		for j := range n {
			fmt.Println("j:", j)
			total := 1
			for i := j; i > 0; i-- {
				total *= i
			}
			out <- total
			close(out)
		}
	}(c, n)

	return out
}
