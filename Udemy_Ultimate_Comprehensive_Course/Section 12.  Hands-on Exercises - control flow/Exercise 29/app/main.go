package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for j := 0; j < 100; j++ {
		fmt.Printf("\tj =\t%v\n", j)
	}

	for i := 1; i <= 100; i++ {
		fmt.Printf("\n\n----- Iteration %v -----\n", i)

		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x = %v, y = %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Printf("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Printf("x is greater than or equal to 4 and less than or equal to 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Printf("none of the previous cases were met")
		}

	}
}
