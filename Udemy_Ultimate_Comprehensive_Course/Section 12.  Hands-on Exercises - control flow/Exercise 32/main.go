package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		if i > 50 {
			break
		}
		fmt.Printf("i = %v\n", i)
		i++
	}
}
