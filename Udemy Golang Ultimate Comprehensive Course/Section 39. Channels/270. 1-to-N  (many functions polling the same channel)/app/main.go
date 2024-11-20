package main

import (
	"fmt"
)

func main() {
	n := 10  //nr functions polling from the same channel
	c := make(chan int)
	done := make(chan bool)

	// One function putting values into the channel
	go func() {
		for i := 0; i < 100000; i++ {
			c <-i
		}
		close(c)
	}()


	// Ten functions reading values from the channel
	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
