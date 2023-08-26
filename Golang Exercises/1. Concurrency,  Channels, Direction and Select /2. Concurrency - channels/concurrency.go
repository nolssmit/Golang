package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" 	// channel revecive "ping" message
	}
}

/*
Using a channel like this synchronizes the two goroutines. 
When pinger attempts to send a message on the channel it will wait 
until printer is ready to receive the message. (this is known as blocking) 
Let's add another sender to the program and see what happens. 
*/
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"		// channel revecive "pong" message
	}
}

func printer(c chan string) {
	for {
		msg := <-c 		// channel send "ping" or "pong" message for printing
		fmt.Println(msg)	
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
