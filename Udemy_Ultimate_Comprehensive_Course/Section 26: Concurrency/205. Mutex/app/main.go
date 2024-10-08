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

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	//  https://pkg.go.dev/sync#Mutex
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			// Lock the code below so no one else can access the counter variable
			v := counter
			// time.Sleep(time.Second * 2)
			// Allow something else to run
			runtime.Gosched()
			v++
			counter = v
			// This call of the go function is done
			mu.Unlock()
			//Done decrements the WaitGroup counter by one.
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
