package main

import "fmt"

func main() {

	// array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers!", "How's tody?"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	var cc [3] string
	cc[1] = "xx"
	fmt.Printf("cc=%v of type %T with length %v\n", cc, cc, len(cc))

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)
	// c = a

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
	// Error , because it's outside the block level scope
	// fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)   
}

/*
"Arrays have their place, but they’re a bit inflexible,
so you don’t see them too often in Go code.
Slices, though, are everywhere.
They build on arrays to provide
great power and convenience."
~ Go Slices: usage and internals - Go Blog - Andrew Gerrand
*/
// https://go.dev/blog/slices-intro

// PRACTICE - next video
/*
Use a composite literal array
to store these elements in an array
and let the compiler determine the length of the array
*/