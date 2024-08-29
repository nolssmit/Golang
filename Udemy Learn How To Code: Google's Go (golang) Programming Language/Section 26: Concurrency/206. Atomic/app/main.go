package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GoRoutines at beginning:", runtime.NumGoroutine())
	fmt.Println("------------------------------------")

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Lock the code below so no one else can access the counter variable
			atomic.AddInt64(&counter, 1)
			// Allow something else to run
			runtime.Gosched()
			fmt.Println("Counter in loop:\t", atomic.LoadInt64(&counter))
			// time.Sleep(time.Second * 2)
			wg.Done()
		}()
		fmt.Println("Goroutine run:", runtime.NumGoroutine())
	}
	// Wait till the execution of the go routines
	// are done before exit the program
	wg.Wait()
	fmt.Println("GoRoutines at end:", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
	// Note the "Count" will differ with each run

	// How to trace a race condition in your code:
    // In terminal: go run -race main.go
}
