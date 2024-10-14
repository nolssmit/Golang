package main

import (
	"fmt"
)

func main() {
	q := make(chan int)

	// send
	go sendToChannel(q)

	// receive
	receiveFromChannel(q)

	fmt.Println("about to exit")
}

func sendToChannel(a chan int) <-chan int {
	fmt.Println("about to send")

	defer close(a)  // not needed
	for i := 0; i < 10; i++ {
		a <- i
	}

	return a
}

func receiveFromChannel(b chan int) {
	fmt.Println("about to receive")

	for i := 0; i < 10; i++ {	
		select {
		case v := <-b:
			fmt.Println(v)
		}
	}
}
// Do a data race test:  $ go run --race main.go