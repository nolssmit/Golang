package main

import (
	"fmt"
//	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go updateChannel(i, c)
	}

	select {
	case v := <-c:
		fmt.Println(v)
	}
}

func updateChannel(i int, c chan int) {
//	time.Sleep(1 * time.Second)
	c <- i
}
