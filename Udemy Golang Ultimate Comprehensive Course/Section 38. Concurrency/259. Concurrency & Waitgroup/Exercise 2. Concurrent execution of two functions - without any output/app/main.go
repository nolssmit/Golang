package main

import (
	"fmt"
)

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println(i,"foo")
	}	
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println(i, "bar")
	}	
}
// Note 3 functions running in parallel, not waiting for each other