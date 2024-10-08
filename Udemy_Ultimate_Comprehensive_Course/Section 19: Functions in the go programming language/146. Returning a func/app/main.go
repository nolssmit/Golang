package main

import "fmt"

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int {
	return 42
}

func bar() func() int  {
	return func() int {
		return 43
	}
}

