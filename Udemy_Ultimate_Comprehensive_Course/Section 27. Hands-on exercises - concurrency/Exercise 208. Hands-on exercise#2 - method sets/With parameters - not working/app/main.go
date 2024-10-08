package main

// Pointer Receiver and Pointer Value
import "fmt"

type person struct {
	first string
}

func (p *person) speak(msg string) {
	fmt.Println(msg, p.first)
}

func (p *person) saySomething(msg string) {
	fmt.Println(msg, p.first, "and Goodbye")
}

type human interface {
	speak(msg string)
	saySomething(msg string)
}

func main() {
	p1 := &person{
		first: "Peter",
	}
	p1.speak("Hallo")


	p1.speak("Hi")
	p1.saySomething("Hello")
}
