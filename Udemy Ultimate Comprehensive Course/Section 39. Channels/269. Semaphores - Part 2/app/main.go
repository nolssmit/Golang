package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)
	
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()	
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	// wait to receive a value from the channel c
	// and print it
	for n := range c {
		fmt.Println(n)
	}
}
