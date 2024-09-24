package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

// Send to channel
	go chanSend(c)

// Receive from channel
	go chanReceive(c) // will print nothing

	chanReceive(c)	  // will print 44

	fmt.Println("About to exit")
}

// Send	to channel
func chanSend(c chan<- int) {
	c <- 44
}

// Receive from channel
func chanReceive(c <-chan int) {
	fmt.Printf("function chanReceive: channel type: %T, channel value: %d\n",c, <-c)
}