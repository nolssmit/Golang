package main

import (
	"fmt"
)

func addF(a,b float64) float64 {
	return a + b
}

func addI(a, b int) int {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

func main(){
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.3))
	fmt.Println("--- calling a generic function ---")
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.3))	
}