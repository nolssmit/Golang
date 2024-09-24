package main

import (
	"fmt"
)

func main() {
	q := make(chan int)

	// send
	c := sendToChannel(q)

	// receive
	receiveFromChannel(c, q)

	fmt.Println("about to exit")
}

func sendToChannel(q chan<- int) <-chan int {
	c := make(chan int)

	fmt.Println("about to send")
	go func(){
		defer close(c)  // perhaps better way to close
		defer close(q)	// both channels with defer
		for i := 0; i < 10; i++ {	
			c <- i
		}
		q <- 1
	//	close(c)
	}()

	return c
}


func receiveFromChannel(c, q <-chan int) {
	fmt.Println("about to receive")

	for {
		select {
		case v := <- c:
			fmt.Println(v)
		case <- q:
			return
		}
		
	}
}
