package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
// wg has package scope

// run once to set things up before func main() runs
func init() {

}

func main() {
	// Print some constants
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// Add 1 to the WaitGroup varible: There is one thing to wait for
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// program will .Wait() till all the things we are waiting for are done
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// remove one thing we are waiting for and the program 
	// may exit if there are no things to wait for
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
