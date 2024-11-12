package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

// Creating a channel where you can only receive values from
func incrementor() <-chan int {
	out1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out1 <- i
		}
		close(out1)
	}()
	return out1
}

//We take a channel as an argument which you can only receive values from
func puller(c <-chan int) <-chan int {
	out2 := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out2 <- sum
		close(out2)
	}()
	return out2
}
// About channel types, read: https://go101.org/article/channel.html