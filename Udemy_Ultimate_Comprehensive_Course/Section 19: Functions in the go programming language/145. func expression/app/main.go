package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()
	y("Nols")
}

func foo() {
	fmt.Println("Foo ran")
}
