package main

import "fmt"
// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	f := outer()
	fmt.Println(f())

	fmt.Printf("%T\n", f)
}

func outer() func() int  {
	return func() int {
		return 43
	}
}