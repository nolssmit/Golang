package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4,7,3,42,99,18,16,56,12}
	xs :=[]string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("Unsorted:",xi)
	sort.Ints(xi)
	fmt.Println("Sorted:",xi)
	fmt.Println("----------------------")
	fmt.Println("Unsorted:",xs)
	sort.Strings(xs)
	fmt.Println("Sorted:",xs)
}
