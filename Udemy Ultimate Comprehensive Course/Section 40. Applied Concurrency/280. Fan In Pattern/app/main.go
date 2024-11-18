// See: https://go.dev/talks/2012/concurrency.slide#1 
// https://www.youtube.com/watch?v=f6kdp27TYZs&ab_channel=GoogleforDevelopers
// https://vimeo.com/49718712
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		// "<-c" blocks the program from exiting until it receives a value from the channel.
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// Returns receive-only channel
// This is not FAN OUT
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
// Multiple functions writing to the same channel.
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
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