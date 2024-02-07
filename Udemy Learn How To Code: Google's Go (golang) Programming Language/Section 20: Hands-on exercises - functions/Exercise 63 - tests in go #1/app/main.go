package main

import (
	"fmt"
)

func main() {
	fmt.Println("The sum is: ", Add(3, 4))
}

func Add(a int, b int) int {
	return a + b
}