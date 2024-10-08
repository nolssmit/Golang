package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)

			// go func(m int) {
			// 	for i := 0; i < 10; i++ {
			// 		c <- i*10 + m
			// 	}
			// 	wg.Done()
			// }(i)

			// alternative I'm using
			go updateChannel(i, &wg, c)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func updateChannel(m int, wg *sync.WaitGroup, c chan int) {
	for i := 0; i < 10; i++ {
		c <- i*10 + m
	}
	wg.Done()
}