package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("xxx.txt")

	// Different idiomatic ways to handle errors
	if err != nil {
		// fmt.Println("Error happened: ", err)
		log.Println("Error happened: ", err)  // Most populare option
		// log.Fatalln(err)
		// panic(err)
	}
}

/*
 Package log implements a simple logging package ... writes to
 standard error and prints the date and time of each logged message.
 ... the Fatel functions call os.Exit(1) after writing the log
 message  ...  the Panic functions call panic after writing the log
 message.

 log.Println call Output to print the standard logger. Arguments are
 handled in the manner of fmt.Println.

 Package log implements a simple logging package ... writes to
 standard error and prints the data and time of each logged message.

 The Panic functions call panic after writing the log message.
*/
