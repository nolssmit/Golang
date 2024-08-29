package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 1 printing")
}

func print2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function 2 printing")
}

func main() {
	// https://pkg.go.dev/runtime#pkg-functions
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin GoRoutines:", runtime.NumGoroutine())
	fmt.Println("------------------------------------")

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		go print1(wg)
		fmt.Println("Mid CPUs:", runtime.NumCPU())
		fmt.Println("Mid GoRoutines:", runtime.NumGoroutine())
		go print2(wg)
	}()

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
	fmt.Println("------------------------------------")
	fmt.Println("End CPUs:", runtime.NumCPU())
	fmt.Println("End GoRoutines:", runtime.NumGoroutine())
}
