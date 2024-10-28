package acdc

import (
	"fmt"
)

func ExampleSum() {
	fmt.Println(Sum(9, 10))
	// Output: 19
	
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// Output: 55
}