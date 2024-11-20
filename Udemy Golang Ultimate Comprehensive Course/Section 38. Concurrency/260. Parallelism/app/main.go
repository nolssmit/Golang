package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Not required, it's the default option anyway for Golang 1.5+
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println(i,"foo")
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()	
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println(i, "bar")
		time.Sleep(time.Duration(20 * time.Millisecond))
	}	
	wg.Done()
}
// Two processes running concurrently and with visual output