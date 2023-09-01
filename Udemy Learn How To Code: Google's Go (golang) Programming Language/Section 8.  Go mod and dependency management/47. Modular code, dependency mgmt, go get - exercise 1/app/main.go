package main

import (
	"fmt"

	"github.com/nolssmit/puppy"
)

func main() {
	s1 := puppy.Back()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// also like this
	fmt.Println(puppy.Back())
	fmt.Println(puppy.Barks())
}