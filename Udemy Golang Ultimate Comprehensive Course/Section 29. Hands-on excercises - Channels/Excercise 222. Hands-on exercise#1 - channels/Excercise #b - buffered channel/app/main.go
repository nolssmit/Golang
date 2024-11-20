package main

import (
	"fmt"
)

func main() {

	cc := make(chan int, 4)

	cc <- 42
	cc <- 43
	cc <- 44
	cc <- 45

	for i := range cc {
		fmt.Println(i)
	}

}
