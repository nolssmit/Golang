package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines at beginning:", runtime.NumGoroutine())
	fmt.Println("------------------------------------")

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	incrementer := 0
	
	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("Incrementer:", incrementer)
			// Done decrements the WaitGroup counter by one.
			wg.Done()
		}()
		fmt.Println("Goroutine run:", runtime.NumGoroutine())
	}
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()
	fmt.Println("Incrementer end value:", incrementer)
	fmt.Println("GoRoutines at end:", runtime.NumGoroutine())
	fmt.Println("Count:", incrementer)
	// Note the "Count" will differ with each run

	// How to trace a race condition in your code:
    // In terminal: go run -race main.go
}
