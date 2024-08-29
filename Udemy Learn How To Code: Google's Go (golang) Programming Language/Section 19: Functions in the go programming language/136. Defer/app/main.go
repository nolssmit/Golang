package main

import (
	"fmt"
)

func main(){
	// foo() will run after function surroundig it exists (returns).
	// Typical use of a defered function is the closing of files
	defer foo()  

	bar()		 
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo(){
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}

// https://go.dev/ref/spec#Defer_statements