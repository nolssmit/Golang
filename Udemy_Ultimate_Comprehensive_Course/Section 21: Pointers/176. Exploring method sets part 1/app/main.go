package main

import (
	"fmt"
)

type dog struct{
	first string
}

func (d dog) walk(){
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run(){
	d.first = "Laika"
	fmt.Println("My name is", d.first, "and I'm running.")
}


func main(){
	d1 := dog{"Spike"}
	d1.walk()
	d1.run()

	d2 := &dog{"Porthos"}
	d2.walk()
	d2.run()
}