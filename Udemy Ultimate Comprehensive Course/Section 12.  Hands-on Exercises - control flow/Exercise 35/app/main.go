package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("Index %v,  value %v\n", i, v)
	}
}

