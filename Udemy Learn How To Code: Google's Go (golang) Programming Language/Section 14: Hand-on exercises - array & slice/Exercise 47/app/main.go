package main

import (
	"fmt"
)

func main() {
	y1 := make([]int, 50)
	fmt.Printf("Length: %v - Capacity: %v\nValues:%v\n", len(y1), cap(y1), y1)
    fmt.Println("-----------------------------")

	y1[0] = 100
	fmt.Printf("Length: %v - Capacity: %v\nValues:%v\n", len(y1), cap(y1), y1)
    fmt.Println("-----------------------------")

	y1 = append(y1, 50)
	fmt.Printf("Length: %v - Capacity: %v\nValues:%v\n", len(y1), cap(y1), y1)
    fmt.Println("-----------------------------")

	y2 := make([]int, 0, 50)
	fmt.Printf("Length: %v - Capacity: %v\nValues:%v\n", len(y2), cap(y2), y2)
    fmt.Println("-----------------------------")

	y2 = append(y2, 50)
	fmt.Printf("Length: %v - Capacity: %v\nValues:%v\n", len(y2), cap(y2), y2)
    fmt.Println("-----------------------------")

	x1 := make([]string, 0)
	x1 = append(x1, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming",)

	fmt.Printf("Length %v, Capacity %v\n", len(x1), cap(x1))
	for i := 0; i < len(x1); i++ {
		fmt.Printf("Index: %v, - Value: %v\n", i, x1[i])
	}

	x2 := make([]string, 0, 50)
	x2 = append(x2, x1...)

	fmt.Println("-------------------------")

	fmt.Printf("Length %v, Capacity %v\n", len(x2), cap(x2))
	for i := 0; i < len(x2); i++ {
		fmt.Printf("Index: %v, - Value: %v\n", i, x2[i])
	}	
}

/*
For this exercise, do the following:
Create a slice to store the names of all of the states in the United States of America.
Use make and append to do this.
Goal: do not have the array that underlies the slice created more than once.
Print out
the len
the cap
the values, along with their index position, without using the range clause.
Here is a list of the 50 states:

" Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming",


*/
