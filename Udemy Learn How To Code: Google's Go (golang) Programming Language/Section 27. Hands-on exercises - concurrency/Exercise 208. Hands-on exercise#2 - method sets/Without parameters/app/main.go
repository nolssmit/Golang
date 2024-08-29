package main

// Pointer Receiver and Pointer Value
import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Willem",
	}
	p1.speak()
//	saySomething(p1)
	saySomething(&p1)
}
