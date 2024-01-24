package main

import (
	"fmt"
)

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

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

	//	sa1.speak()
	//	p2.speak()

	// polymorphism: a value can be of more than one type
	saySomething(sa1)
	saySomething(p2)
}
