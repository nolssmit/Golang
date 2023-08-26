package main

import (
	"fmt"
	"time"
)

// Package global variable
var (
	count int
)

func main() {
	inteations := 10000

	for i := 0; i < inteations; i++ {
		go increment()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	count++
}