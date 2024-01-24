package main

import "fmt"


// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	foo()
}

func foo() {
	fmt.Println("Foo ran")
}