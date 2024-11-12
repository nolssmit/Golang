package main

import (
	"fmt"
)

// func main() {
// 	c := make(chan int)
// 	c <- 42
// 	fmt.Println(<-c)
// }

// This results in a deadlock.
// Can you determine why?
// And waht would you do to fix it?

//var c = make(chan int)
func main() {
	c := make(chan int)

	go func (c chan<- int) {			
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)	
	}(c)

	// Will show only the first value
	fmt.Println("First value: ", <-c)

	// Receive from the channel, except for the first value
	for v := range(c){
		fmt.Println(v)
	}	
	// Omit the close(c), then you will get a deadlock
	// If you range over the channel then the channel 
	// must be closed after used

	fmt.Println("About to exit")
}

