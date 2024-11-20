package main

import (
	"fmt"
)

func main() {
	jb := [] string {"James", "Bond", "Martini", "Chocolate"}
	jm := [] string {"Jenny", "Bond", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println("-----------------------------------")
	xss := [][] string {jb, jm}
	fmt.Println(xss)
	
}