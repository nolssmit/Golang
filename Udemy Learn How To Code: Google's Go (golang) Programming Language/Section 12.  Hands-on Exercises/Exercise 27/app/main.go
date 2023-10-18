package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("x = %v, y = %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Printf("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Printf("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Printf("none of the previous cases were met")
	}
}
