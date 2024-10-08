package main

import (
	"fmt"
	"github.com/nolssmit/Golang/tree/1317d813aceb72031fca6aca4095a814a3d20791/Udemy%20Learn%20How%20To%20Code%3A%20Google's%20Go%20(golang)%20Programming%20Language/Section%2033.%20Hands-on%20excercises%20-%20Documentation/Exercise%201.%20create%20dog%20package%20to%20convert%20dog%20years%20to%20human%20years/app/main.go"
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