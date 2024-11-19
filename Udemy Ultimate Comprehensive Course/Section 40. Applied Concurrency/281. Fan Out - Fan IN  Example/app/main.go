package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2,3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
   	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels
	for n := range merge(c1,c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
	
// passing a variadic parameter and what you are getting is a slice
func gen(nums ...int) chan int {
	fmt.Printf("Type of nums is %T\n", nums)

	out := make(chan int)
	go func() {
		// range over a slice getting ignoring the key but getting the value
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		// range over a channel
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	fmt.Printf("Type of cs is %T\n", cs)

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))


	for _, c := range cs {
		// We pass the parameter c to limit the scope of the anonymous goroutine
		// to cs's variables one at a time.
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}
/*
code source
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25
*/

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed.

FAN IN
A function can read from multiple inputs and process until all are closed by
multiplexing the inputs into a single channel that's closed when
all inputs are closed.

PATTERN
There's a pattern to our pipeline functions:
--- stages close their outbound channels when all the send operations are done.
--- stages keep receiving values from inbound channels until those channels are closed.

source:
https://blog.golang.org/pipelines
*/

/* 
CHALLENGE #1
When know HOW to do FAN OUT and FAN IN, but do we know WHY?
Why would we want to do fan out / fan in?
*/