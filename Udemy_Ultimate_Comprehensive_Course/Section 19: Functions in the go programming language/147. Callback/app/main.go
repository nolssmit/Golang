package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main() {
	fmt.Printf("%T,\n", add)
	fmt.Printf("%T,\n", subtract)
	fmt.Printf("%T,\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a+b
}

func subtract(a,b int) int {
	return a-b
}