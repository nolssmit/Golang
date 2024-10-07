package main

import (
	"fmt"

	"Udemy Learn How To Code: Google's Go (golang) Programming Language/Section 33. Hands-on excercises - Documentation/Exercise 1. create dog package to convert dog years to human years/app/dog.go"
)

type canine struct {
	name string
	age  dog.Years
}

func main() {
	fido := canine{
		name: "Fido",
		age:  4,	
	}
	fmt.Println(fido)
}