package main

import (
	"fmt"
	"github.com/nolssmit/Udemy_Ultimate_Comprehensive_Course/Section_33_Hands-on-excercises-Documentation/Exercise_1_convert-dog-years-to-human-years/app/dog"
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