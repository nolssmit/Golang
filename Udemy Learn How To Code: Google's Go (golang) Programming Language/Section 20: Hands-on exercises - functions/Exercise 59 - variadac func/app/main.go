package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	xi := []int {1,2,3,4,5,6,7,8,9}
	fmt.Println("The sum is", foo(xi...))

	fmt.Println("The sum is", bar(xi))
}

func foo(ii ...int) int {
	t := 0
	for _, val := range ii{
		t += val
	}
	return t
}

func bar(ii[] int) int {
	t := 0
	for _, val := range ii{
		t += val
	}
	return t
}