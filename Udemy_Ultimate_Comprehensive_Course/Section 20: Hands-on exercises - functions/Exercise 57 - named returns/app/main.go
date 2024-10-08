package main

import (
	"fmt"
)

// a named function with an identifier
// func (r receiver) identifier(p parameter(s)) (r return(s)) { code } (arguments)

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }

func main() {
	x := sum([]int{1,2,3,4,5,6,7,8,9})
	fmt.Println(x)
}

// Todd does't like name returns
/*
func sum(ii [] int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}
*/

func sum(ii [] int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}