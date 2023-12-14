package main

import "fmt"

func main() {
	var a[5]int
	for i := 0; i < 4; i++ {
		a[i] = i
	}

	fmt.Println(a)
	// Note last element has default value of zero

	b := [2]string{"Penn", "Teller"}
	for i := 0; i <= 1; i++ {
		fmt.Println(b[i])
	}	
	fmt.Println(b)
}
