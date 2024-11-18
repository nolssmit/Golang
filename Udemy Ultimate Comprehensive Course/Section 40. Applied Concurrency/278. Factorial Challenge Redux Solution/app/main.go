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
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
