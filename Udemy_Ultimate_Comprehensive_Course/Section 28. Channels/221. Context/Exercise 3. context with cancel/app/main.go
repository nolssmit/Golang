package main

import (
	"context"
	"fmt"
	"runtime"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	fmt.Println("error check 1: ", ctx.Err())
	fmt.Println("num goroutines 1: ", runtime.NumGoroutine())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return  // return not to leak the routine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}