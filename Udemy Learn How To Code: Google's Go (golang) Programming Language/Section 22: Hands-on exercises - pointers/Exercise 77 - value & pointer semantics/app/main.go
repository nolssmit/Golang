package main

import (
	"fmt"
)

type person struct{
	first string
}

func main(){
	p := person{
		first: "Jenny",
	}
	fmt.Println(p)
	p = changeName(p, "Jen")
	fmt.Println(p)

	changeNamePrt(&p, "JMo")
	fmt.Println(p)
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePrt(p *person, s string){
	p.first = s
}