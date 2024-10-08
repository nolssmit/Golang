package main

import "fmt"

func main(){
	x := 42
	fmt.Printf("%v\t%T\t\n", &x, &x)
	s := "Nols"
	fmt.Printf("%v\t%T\t\n", &s, &s)
}