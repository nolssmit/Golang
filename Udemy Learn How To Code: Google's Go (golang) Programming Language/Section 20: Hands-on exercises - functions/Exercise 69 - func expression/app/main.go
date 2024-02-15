package main

import ("fmt")


// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main(){
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	z()

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 45
	}
}

var z = func(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}