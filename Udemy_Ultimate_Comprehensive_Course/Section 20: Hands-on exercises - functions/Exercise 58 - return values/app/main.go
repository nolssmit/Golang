package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	x := foo(10)
	fmt.Println(x)

	i, s := bar(20, "x")
	fmt.Println(i,s)
}

func foo(x int) int {
	return 2*x
}

func bar (x int, s string) (int, string) {
	return 2*x, s+s
}
