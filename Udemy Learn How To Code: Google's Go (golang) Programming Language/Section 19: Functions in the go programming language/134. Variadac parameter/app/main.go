package main

import (
	"fmt"
)

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is:", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

// The function's final incomming parameter, a variadic parameter is a slice of int
// .. and unlimited number of arguments of a certain type
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

// https://go.dev/ref/spec#Keywords
// See: "Passing arguments to ... parameters"