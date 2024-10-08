package main

import (
	"fmt"
)

func receiveFromChan(cr <-chan int) int {
	 return <- cr
}

func sendToChan(i int, cs chan<- int) {
	cs <- i
}

func main() {
	// Faulty code	
	//cs := make(chan<- int)
	//cs := make(<-chan int)

	cs := make(chan int)
	go sendToChan(42, cs)
	result := receiveFromChan(cs)

	fmt.Printf("Send channel ch type: \t%T\n", cs)
	fmt.Println("Received channel value: ", result)
}
