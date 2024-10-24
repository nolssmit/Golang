package main

import (
	"fmt"
)

func main() {
	// Create a new channel
	c := make(chan int)
	go func() {
		// Send a value into the channel
		// This will block until the receiver gets it
		c <- 42
	}()
	// Read the value from the channel
	// This will block until the sender sends it
	fmt.Println(<-c)
}
//https://pkg.go.dev/github.com/eapache/channels