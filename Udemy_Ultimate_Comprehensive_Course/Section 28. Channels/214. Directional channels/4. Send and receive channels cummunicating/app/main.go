package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) // bidirectional (from and to outside)
	cr := make(<-chan int) // receive-only (from outside)
	cs := make(chan<- int) // send-only (to outside)
	c <- 2
	c <- 3
	fmt.Println("Received from channel c:", <-c)
	fmt.Println("Received from channel c:", <-c)
	fmt.Println("--------------------------")

	// cr <- 4 	//invalid operation: cannot send to receive-only channel cr
	// fmt.Println(<- cs) // invalid operation: cannot receive from send-only channel cs

	fmt.Printf("Bidirectional channel: c\t%T\n", c)
	fmt.Printf("Receive-only channel: cr\t%T\n", cr)
	fmt.Printf("Send-only channel: cs\t%T\n", cs)
	fmt.Println("--------------------------")

	// moving from a specific to general channel not allowed
	// c = cr // cannot use cr (variable of type <-chan int) as chan int value in assignment
	// c = cs // cannot use cs (variable of type chan<- int) as chan int value in assignment

	// moving from a general to specific channel is allowed
	fmt.Printf("Channel c now receice only: c\t%T\n", (<-chan int)(c))
	fmt.Printf("Channel c now send only: c\t%T\n", (chan<- int)(c))
}
