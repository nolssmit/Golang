package main

// Pointer Receiver and Pointer Value
import "fmt"

type person struct {
	first string
	message string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.message)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := &person{
		first: "Willem",
		message: "Hello",
	}
	p1.message = p1.message + " and " + "Goodbye"
	p1.speak()

	var h human = p1
	h.speak()

}
