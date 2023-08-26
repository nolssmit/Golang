package main

import (
	"fmt"
	"sync"
	"time"
)

// Package global variable
var (
	lock sync.Mutex
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
	lock.Lock()
	// Only one go routine has access to the global package variable count
	count++
	lock.Unlock()
}