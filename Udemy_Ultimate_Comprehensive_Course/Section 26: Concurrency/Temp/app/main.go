package main

import (
	"fmt"
	"sync"
	"runtime"
)


 
type message struct {
	aMessage string
}

type booskap interface {
	stuur() string
}

func (wg *sync.WaitGroup ) stuur(name string) string {
	defer wg.Done()
	fmt.Println("----------------------------------")
	return ("Hello " + name)

}

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin GoRoutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		  go stuur("Koos")
	}()

	wg.Wait()

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Begin CPUs:", runtime.NumCPU())
	fmt.Println("Begin GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Klaar")
}

func stuur(s string) {
	panic("unimplemented")
}
