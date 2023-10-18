package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("Iterration %v,\t random number = 0\n", i+1)
		case 1:
			fmt.Printf("Iterration %v,\t random number = 1\n", i+1)
		case 2:
			fmt.Printf("Iterration %v,\t random number = 2\n", i+1)
		case 3:
			fmt.Printf("Iterration %v,\t random number = 3\n", i+1)
		case 4:
			fmt.Printf("Iterration %v,\t random number = 4\n", i+1)
		}
	}
}
