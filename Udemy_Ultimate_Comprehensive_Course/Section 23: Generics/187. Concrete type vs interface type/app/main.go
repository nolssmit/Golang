package main

import "fmt"

// Concrete type
type Employee struct {
	Name string
	Age int
	}

// Interface type
type Reader interface {
	Read(p []byte) (n int, err error)
	}


func main(){
	emp := Employee{Name: "John Doe", Age: 25}
	fmt.Println(emp)
}
