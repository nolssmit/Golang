package main

import (
	"fmt"
)

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
		for i := 0; i < 100; i++ {
			for j := 1; j < 10; j++ {
				out <-j
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
/*
CHALLENGE #1
-- Change the code above to execute 1000 factorial computations concurrently and in in parrallel.
-- Use the "fan out / fan in" pattern to accomplish this.

CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
*/