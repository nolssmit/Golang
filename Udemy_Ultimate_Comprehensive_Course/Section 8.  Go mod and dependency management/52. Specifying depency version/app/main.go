//https://www.digitalocean.com/community/tutorials/how-to-use-go-modules
package main

import (
	"fmt"

	"github.com/nolssmit/Golang/GolangPackages/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)	

	// also like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
}