package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	p1 := struct {
		first string
		last  string
		age   int
		// ltk   bool   will give an compile error
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		age:   27,
	}

	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
	fmt.Printf("%#vT \n", p2)

	p2 = p1
	fmt.Printf("%T \n", p2)
	fmt.Printf("%#vT \n", p2)
}
