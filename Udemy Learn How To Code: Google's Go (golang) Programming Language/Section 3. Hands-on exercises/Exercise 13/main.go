package main

import "fmt"

/*
write a program that declares two variables
	● one variable to store a VALUE of TYPE int8
		○ assign to it the largest number possible, then print it
	● one variable to store a VALUE of TYPE uint8
		○ assign to it the largest number possible, then print it
*/

func main() {
	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T\n", x, x)
	fmt.Printf("%v is of type %T\n", y, y)
}