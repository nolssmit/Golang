package main

import (
	"fmt"
)

func main() {
	a := 1
	fmt.Println((a))  // run with this line
//	fmt.Println((&a)) // run with this line 
}
// In terminal, run with: go run -gcflags -m main.go


/*
escapes to the heap
variable shared between main() and Println()

moved to heap
variable moved to the heap
*/
