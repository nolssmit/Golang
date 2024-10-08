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

	func(){
		fmt.Println("Anonymous func ran")
	}()

	func(s string){
		fmt.Println("This is an anonymous func showing my name", s)
	}("Nols")
}

func foo(){
	fmt.Println("Foo ran")
}