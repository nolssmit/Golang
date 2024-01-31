package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

type Person struct{
	first string
	age int
}

func (p Person) Speak() {
	fmt.Println("I'm", p.first, "and I'm", p.age, "years old")
}
func main(){
	p1 := Person{
		first: "Nols",
		age: 73,
	}

	p1.Speak()
}