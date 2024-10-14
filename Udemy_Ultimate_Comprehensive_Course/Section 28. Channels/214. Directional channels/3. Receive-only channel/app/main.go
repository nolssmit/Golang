package main

import (
	"fmt"
)

func main() {
	// Create a new buffered receive-only (from) channel that hold up to 2 values.
	c := make(<- chan int, 2)
	fmt.Printf("%T\n", c)

	// Can't send values to the channel
	c <- 42
	c <- 43

	 fmt.Println(<-c)
	 fmt.Println(<-c)
	 fmt.Println("---------")
}
// https://go.dev/doc/effective_go#channels
// https://pkg.go.dev/github.com/eapache/channels
// https://go-proverbs.github.io/