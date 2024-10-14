package main

import (
	"fmt"

)

// func populate() {
// 	c := make(chan int, 1)

// 	go func() {
// 		c <- 42
// 		c <- 43
// 	}()

// 	fmt.Println(<-c)
// 	fmt.Println(<-c)
// }

func main() {

	c := make(chan int)

	go func(cc chan int) {
		cc <- 42
		cc <- 43
		cc <- 44
		cc <- 45
		close(cc)
	}(c)


	for i :=  range c {
		fmt.Println(i)
	}

}
