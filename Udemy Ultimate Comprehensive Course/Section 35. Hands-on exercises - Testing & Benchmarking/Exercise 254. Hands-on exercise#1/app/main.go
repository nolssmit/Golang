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
	fmt.Println(dog.Years(20))
	dog.
	fmt.Println(dog.YearsTwo(20))
}