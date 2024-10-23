package main

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/dog"
)
type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),	
	}
	fmt.Println(fido)


	fmt.Printf("%s is %d years old\n", fido.name, dog.Years(fido.age))
}