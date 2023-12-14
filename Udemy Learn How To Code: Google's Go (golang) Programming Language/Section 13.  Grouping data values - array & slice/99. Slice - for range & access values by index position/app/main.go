package main

import (
	"fmt"
)

func main(){
	xs := [] string {"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Printf("xs: %v\n", xs)

    fmt.Println("-----------------")

	for _, v := range xs {
		fmt.Printf("value: %v\n", v)
	}

    fmt.Println("-----------------")

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

    fmt.Println("-----------------")

	for i :=0; i<len(xs); i++ {
		fmt.Println(xs[i])
	}

	fmt.Println("-----------------")

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
//	fmt.Println(xs[3])   ... does not work

fmt.Println("-----------------")

    fmt.Println(len(xs))

	fmt.Println("-----------------")

	for i :=0; i<len(xs); i++ {
		fmt.Println(xs[i])
	}

	fmt.Println("-----------------")

	xi := [] int{1, 2, 3, 4, 5}	
	fmt.Println(len(xi))

	fmt.Println("-----------------")

	fmt.Println(xi[0])	
	fmt.Println(xi[1])	
	fmt.Println(xi[2])	
	fmt.Println(xi[3])	
	fmt.Println(xi[4])			

	fmt.Println("-----------------")

	for i, v := range xi {
		fmt.Println(i, "-", v)
	}
}