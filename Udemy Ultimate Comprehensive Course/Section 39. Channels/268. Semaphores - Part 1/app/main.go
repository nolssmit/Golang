package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// we block here until done <- true
	go func() {
		// listens for done deal, then closes the channel
		<-done
		<-done
		close(c)
	}()
	
	// to unblock above we need to take values of chan c here
	// but we never get here, because we're blocked above
	for n := range c {
		fmt.Println(n)
	}
}
