package main

import (
//	"fmt"
//	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		// fmt.Println("err happened: ", err)  // to std-out
		// fmt.Println("----------------------") 
		// log.Println("err happened", err) // default to std-out with a date time stamp but we may write to a file
		// log.Fatalln("err happened", err)  // os.Exit(1)
		// log.Panicln("err happened", err) // defered function run & can use "recover"
		// panic(err)   
	}
}
// Println format using default formats for its operands and writes to standard output.