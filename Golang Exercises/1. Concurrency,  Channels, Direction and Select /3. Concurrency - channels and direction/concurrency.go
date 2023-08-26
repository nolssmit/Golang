package main

import (
	"fmt"
	"time"
)

/*
Using a channel like this synchronizes the two goroutines.
When pinger attempts to send a message on the channel it will wait
until printer is ready to receive the message. (this is known as blocking)
Let's add another sender to the program and see what happens.
*/


func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		for {
			c1 <- "Hello from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "Hello from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1: ", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2: ", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
