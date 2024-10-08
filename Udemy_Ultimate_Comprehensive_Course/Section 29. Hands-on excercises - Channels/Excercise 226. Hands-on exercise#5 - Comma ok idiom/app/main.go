package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		//	defer close(c)
		c <- 42
	}()

	v, ok := <-c	// declare v and ok
	fmt.Println("using comma ok: channel value received, value: ", v)

	close(c)

	v, ok = <-c		// assign v and ok
	fmt.Println("using comma ok: channel value not received:", v, ok)

	// receive(c)
}

// func receive(c <-chan int) {
// 	select {
// 	case v, ok := <- c:
// 		if !ok {
// 			fmt.Println("using comma ok: channel value not received", v, ok)
// 			return
// 		} else {
// 			fmt.Println("using comma ok: channel value received:, value: ", v)
// 		}
// 	}
// }
