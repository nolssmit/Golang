package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Println(a)
	a++
	fmt.Println(a)

	m := make(map[string] int)
	fmt.Println(m["beatiful"])
	m["beatiful"]++
	fmt.Println(m["beatiful"])	
	m["beatiful"]++
	fmt.Println(m["beatiful"])	
}