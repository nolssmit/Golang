package main

import "fmt"


var y int


/*
write a program that uses the following:
	● var for zero value
	● short declaration operator
	● multiple initializations
	● var when specificity is required
	● blank identifier
print items as necessary to make the program interesting
*/

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 43, "hapiness"
	fmt.Println(a, b)

	var c float32 = 42.777
	fmt.Printf("%v of type %T\n", c, c)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)
}