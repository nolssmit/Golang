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

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var incrementer int64
	incrementer = 0

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()
			fmt.Println("Incrementer:", atomic.LoadInt64(&incrementer))

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
}
