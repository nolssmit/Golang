package main

import (
	"fmt"
)

func main() {
	sx := [] string {"1", "2a", "3"}
	fmt.Println(sx)
	fmt.Println("----------------------------")
	yx := make([]int, 0, 10)
	fmt.Println(yx)
	fmt.Println(len(yx))
	fmt.Println(cap(yx))
	fmt.Println("----------------------------")
	yx = append(yx, 0,1,2,3,4,5,6,7,8,9,)
	fmt.Println(yx)
	fmt.Println("----------------------------")
	yx = append(yx, 10,11,12,13)
	fmt.Println(yx)
	fmt.Println(len(yx))
	fmt.Println(cap(yx))
	fmt.Println("----------------------------")
}