package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
/*	
	----- COMMAND LINE ARGUMENTS -----
	You can pass values to your program
	  from the command line, e.g.

	Create gotut-13.go using VS Code as in the example

	Open a command line window in the 
	  folder containing the files.
	 $ go build gotut-13.go

	To run the compiled program in the command line window   
	  $ ./gotut-13 24 43 12 9 10

	The program returns an array with everything
	  passed with the name of the app in the first index
	  and also outputs the max number passed in
*/
	// Print the arguments passed
	fmt.Println(os.Args)
	// Get all values starting from the first index till the end
	args := os.Args[1:]

	// Create int array from string array
	iArgs := []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	// Find the maximum value passed in the arguments
	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value :", max)
}