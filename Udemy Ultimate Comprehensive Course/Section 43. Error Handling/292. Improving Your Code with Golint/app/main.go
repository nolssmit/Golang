package main

import (
	"fmt"
)

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	// if n > 10 {
	// 	return fmt.Sprint("x is greater than 10")
	// } else {
	// 	return fmt.Sprint("x is less than 10")
	// }

	// After running golint andfixing the style

	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	}

	return fmt.Sprint("x is less than 10")
}

// install with: $ sudo apt install golint
// Use it with: $ golint
