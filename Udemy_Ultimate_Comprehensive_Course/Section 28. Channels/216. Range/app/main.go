package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	// Send
	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
		close(c)
	}(c)

	// Receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
