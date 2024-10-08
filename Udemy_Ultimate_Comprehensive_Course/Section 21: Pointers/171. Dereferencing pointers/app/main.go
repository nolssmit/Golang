package main

import "fmt"

func main(){
	x := 42
	fmt.Printf("%v\t%T\t\n", x, &x)

	y := &x
	fmt.Printf("%v\t%T\t\n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 43
	fmt.Println(x)
	fmt.Println(*y)
}