package main

import (
	"fmt"
)

func main() {
	fmt.Println("sum of 1 to 3 = ", mySum(1, 2, 3))
	fmt.Println("sum of 1 to 7 = ", mySum(1, 2, 3, 4, 5, 6, 7))
	fmt.Println("sum of 1 to 10 = ", mySum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum  // the test will fail if you return sum + 1
	}