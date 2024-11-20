package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //unbuffered channel, takes only one value
		go func() {
			for i := 0; i < 10; i++ {
				c <- i  
				// the go routine puts into the channel and stops
			}
		}()

		go func() {
			for {
				fmt.Println(<-c)  
				// the go routine reads from the channel and stops
			}
		}()

	time.Sleep(time.Second)	
	// to give ample time to the go routines to execute
}
// Check for a race condition: $ go run -race main.go