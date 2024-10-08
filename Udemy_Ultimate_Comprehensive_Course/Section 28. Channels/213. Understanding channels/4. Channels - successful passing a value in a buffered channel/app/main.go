package main

import (
	"fmt"
)

func main() {
	// Create a new buffered channel with a capacity of 1 value
	c := make(chan int, 2)

	// Send a value into the channel
	// This will block until the receiver gets it
	c <- 42
	// Send another value into the channel
	// but cannot block because the channel is full
	c <- 43

	// Read the value from the channel
	// This will block until the sender sends it
	fmt.Println(<-c)
	fmt.Println(<-c)
}
// https://go.dev/doc/effective_go#channels
// https://pkg.go.dev/github.com/eapache/channels
// https://go-proverbs.github.io/