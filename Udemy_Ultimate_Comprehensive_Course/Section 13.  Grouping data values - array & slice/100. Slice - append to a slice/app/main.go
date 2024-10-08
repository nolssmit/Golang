package main

import (
	"fmt"
)

func main(){
	xi := [] int {42, 43, 44}
	fmt.Println(xi)

	fmt.Println("-------------------")

	// variadic parameter
	xi = append(xi, 45, 46, 47)

	fmt.Println(xi)

	fmt.Println("-------------------")

	yi := [] int {200, 300, 400}
	xi = append(xi, yi...)

	fmt.Println(xi)
}