// https://www.udemy.com/course/go-language/learn/lecture/35216238#notes
//
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make and start off with six elements
	sl1 := make([]string, 6)

  // Assign values by index
	sl1[0] = "Society"
	sl1[1] = "0f"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
  pl("slice sl1 with 5 elements", sl1)

  // Size of slice
	pl("Slice Size :", len(sl1))

  // Cycle with for
  for i := 0; i < len(sl1); i++ {
    pl(sl1[i])
  }

  // Cycle with range
  for _, x := range sl1{
    pl(x)
  }

  // Create a slice literal
  sl2 := []int{22, 12, 2022}
  pl(sl2)

 	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array 
  sArr := [7]int{1, 2, 3, 4, 5, 6, 7}
  pl("The array we are going to slice: ", sArr)
  sl3 := sArr[:3]  // not including the 2'nd
  pl("First three slice : ", sl3)
  sl33 := sArr[3:]
  pl("Slice from 3'rd till last: ", sl33)

  // Get slice from beginning
  sl4 := sArr[:3]
  pl("From first index up to, but not including 3'rd :",sl4)

  // Get slice from third index to the end
  sl5 := sArr[3:]
  pl("From 3'rd index till end, and including 3'rd :",sl5)  

  // Changing the array also changes the slice
   sArr[0] = 10
   pl("sl3, where sl3 := sArr[0:2] :",sl3) 

	// Changing the slice also changes the array
	sl3[0] = 100
	pl("sArr where sl3[0] = 100:", sArr)   

  // Append a value to a slice (Also overwrites array) 
  sl3 = append(sl3, 12)
  pl("sl3, with 12 appended :",sl3)
  pl("sArr, with sArr[0] = 10, replacing original value :", sArr)

  // Printing empty slices will return nulls which show as empty slices
  sl6 := make([]string, 6)
  pl("sl6 :", sl6)
  pl("sl6[0] :", sl6[0])
}
