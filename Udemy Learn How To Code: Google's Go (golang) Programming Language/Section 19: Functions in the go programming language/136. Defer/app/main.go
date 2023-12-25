package main

import (
	"fmt"
)

func main(){
	defer foo()  // will run after function surroundig it exists (returns)
	bar()		 // typical use is the closing of files
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo(){
	fmt.Println("foo")
}

func bar(){
	fmt.Println("bar")
}
