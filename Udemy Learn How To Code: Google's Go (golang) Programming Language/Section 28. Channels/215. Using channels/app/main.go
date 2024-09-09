package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

// Send

	go foo(c)
// Receive
	go bar(c) // will print nothing
	bar(c)

	fmt.Println("About to exit")
}

// Send	to
func foo(c chan<- int) {
	c <- 44
}
// Receive from
func bar(c <-chan int) {

	fmt.Printf("function bar: channel type: %T, channel value: %d\n",c, <-c)
}