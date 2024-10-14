package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	go gen(c)

	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan int) <-chan int {
	defer close(c)

	for i := 0; i < 10; i++ {
		c <- i
	}

	return c
}


func receive(c chan int) {

	fmt.Println("about to receive")
	for v := range c {
		fmt.Println(v)
	}
}
