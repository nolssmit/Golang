package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	fmt.Println("------range------------")
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("--------select---------")
	for i := 0; i < 10; i++ {
		select {
		case v := <-c:
			fmt.Println(v)
		}
	}

	receiveFromChannel(c)
	fmt.Println("---------done-----------")
}

func receiveFromChannel(c chan int) {
	fmt.Println("---------range-----------")
	for v := range c {
		fmt.Println(v)
	}
}
