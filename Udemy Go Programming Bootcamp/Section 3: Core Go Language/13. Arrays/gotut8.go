// https://www.udemy.com/course/go-language/learn/lecture/35216226#notes
//
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
  // ----- ARRAYS -----
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	// Declare integer array with 5 elements
  var arr1 [5]int

  // Assign value to index
  arr1[0] = 1

  // Declare and initialize
  arr2 := [5]int{1,2,3,4,5}
  pl("arr2: ", arr2)

  // Get by index
  pl("Index 0 of arr2 :",arr2[0])

  // Length
  pl("Array Length of arr2 :", len(arr2))
  
  //  Iterate with index
  for i :=0; i < len(arr2); i++ {
    pl(arr2[i])
  }

  // Iterate with range
  fmt.Printf("\nArray elements\n")
  for i, v := range arr2 {
    fmt.Printf("%d : %d\n", i, v)
  }

  // Multidimensional Array
  pl("\n-------Multi dimentional arrays-------\n")  
  arr3 := [2][3]int {
    {1,2,3},
    {4,5,6},
  }

  // Print multidimensional array
  for  i := 0;  i < 2;  i++ {
    for j := 0; j < 3; j++ {
      pl(arr3[i][j])
    }
  }


  // Slice a string into runes
  pl("\n--------String to runes (decimal value of character)-------\n")
  aStr1 := "aceg"
  rArr := []rune(aStr1)
  for _, v := range rArr{
    fmt.Printf("Rune Array : %d\n", v)
  }

  // Convert a byte array to string
  pl("--------Byte array to string--------")
  byteArr := []byte{'a','b','c'}
  bStr := string(byteArr[:])
  pl("I'm a string :", bStr)
}
