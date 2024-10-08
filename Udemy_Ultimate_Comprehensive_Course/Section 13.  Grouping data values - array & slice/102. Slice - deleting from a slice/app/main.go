package main

import (
	"fmt"
)

func main() {
	xi := [] int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-----------------------")
	yi := append(xi[:4], xi[5:]...)
	fmt.Printf("yi - %#v\n", yi)
	fmt.Println("-----------------------")
	zi := append(xi, 100, 200, 300)
	fmt.Printf("zi - %#v\n", zi)
	fmt.Println("-----------------------")
	wi := append(xi, zi...)
	fmt.Printf("wi - %#v\n", wi)
	fmt.Println("-----------------------")	
	ui := append(xi[:2], zi[4:]...)
	fmt.Printf("ui - %#v\n", ui)
	fmt.Println("-----------------------")		
}