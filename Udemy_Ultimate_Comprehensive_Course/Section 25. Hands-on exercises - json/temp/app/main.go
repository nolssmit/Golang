package main

import (
	"fmt"
)

func myFunction(x int, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	fmt.Println("Sum:", sum, "Product:", product)
	
	return
}

func main() {
	fmt.Println(myFunction(5, 2))
}
