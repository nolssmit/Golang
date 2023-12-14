package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Outer loop counter i =%v\n", i)
		j := 0
		for {
			if j == 5 {
				break;
			}
			fmt.Printf("\tInner loop counter j = %v\n", j)
			j++
		}
	}
}