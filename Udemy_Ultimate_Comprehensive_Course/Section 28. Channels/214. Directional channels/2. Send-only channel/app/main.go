package main

import (
	"fmt"
)

func main() {
	// Create a new buffered send-only (to) channel that can hold up to 2 values.
	c := make(chan <- int, 2)
	fmt.Printf("%T\n", c)

	c <- 42
	c <- 43

	// Can't receive values from channel
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("---------")
}

// https://go.dev/doc/effective_go#channels
// https://pkg.go.dev/github.com/eapache/channels
// https://go-proverbs.github.io/
