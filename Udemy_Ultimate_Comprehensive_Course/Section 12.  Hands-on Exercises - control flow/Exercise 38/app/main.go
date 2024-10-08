package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key: %v, \t value: %v\n", k, v)
	}

	fmt.Println("----------------------")

	age1 := m["James"]
	fmt.Println("The age of Bond:", age1)

	fmt.Println("----------------------")

	if v, ok := m["James"]; ok {
		fmt.Println("There is a Bond lookup entry, and Bond's age:", v)
	}

	fmt.Println("----------------------")

	age2 := m["Q"]
	fmt.Println("The age of Q:", age2)

	fmt.Println("----------------------")

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is zero value of int:", v)
	}

	fmt.Println("----------------------")

	c := 1
	for i := 0; i < 100; i++ {
		fmt.Printf("Iteration %v\n", i)
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("\t count: %v, x = %v\n", c, x)
			c++
		}
	}
}
