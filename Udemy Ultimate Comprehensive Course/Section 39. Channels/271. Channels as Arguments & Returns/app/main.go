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

func incrementor() chan int {
	out1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out1 <- i
		}
		close(out1)
	}()
	return out1
}

func puller(c chan int) chan int {
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