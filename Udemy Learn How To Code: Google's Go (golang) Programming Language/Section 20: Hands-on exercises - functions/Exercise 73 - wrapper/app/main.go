package main

import (
	"fmt"
	"time"
)

func main() {
	// Call the wrapped function
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func timeFunc(f func()){
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

/*
To summarize, a wrapper function wraps or modifies another function's behavior, while a
callback function is a function passed as an argument to be executed at a specific point or
event.
*/