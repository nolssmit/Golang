package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("Hello, World!") 

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)  // number of bytes written, including new line
}