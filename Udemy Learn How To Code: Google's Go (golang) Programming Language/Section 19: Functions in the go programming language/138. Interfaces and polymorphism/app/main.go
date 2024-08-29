package main

import (
	"fmt"
)

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

// Also of type human
type person struct {
	first string
}

// Also of type human
type secretAgent struct {
	person
//	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

// Any value of a certain type that has 
// the method speak() will also be of this type
// Hey baby, if you got these methods then you are my type
// Polymorphism is the ability of a VALUE of a certain TYPE to also be of another TYPE.
// An interface in Go defines a set of method signatures.
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
	}

	p2 := person{
		first: "Jenny",
	}

	fmt.Println("------------- Illustrating interfaces -------------")
	sa1.speak()
	p2.speak()
	fmt.Println("------------ Illustrating polymorphism ------------")
	// polymorphism: a value can be of more than one type
	saySomething(sa1)
	saySomething(p2)
}
