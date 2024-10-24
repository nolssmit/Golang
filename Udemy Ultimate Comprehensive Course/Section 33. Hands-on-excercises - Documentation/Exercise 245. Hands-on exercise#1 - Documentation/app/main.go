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

	fmt.Printf("%s is %d years old\n", fido.name, fido.age)
}

// To view documentation: $ godoc -http=localhost:8080
// Then: http://localhost:8080/pkg/github.com/nolssmit/Golang/GolangPackages/dog/
