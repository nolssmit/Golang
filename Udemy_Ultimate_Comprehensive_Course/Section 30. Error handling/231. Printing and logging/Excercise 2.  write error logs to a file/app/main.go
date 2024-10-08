package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("err happened: ", err)
	}
	defer f.Close()
	log.SetOutput(f)


	f2, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("fmt.Println()------------------------")
		// fmt.Println("err happened: ", err)
		// fmt.Println("log.Println()------------------------")
		// log.Println("err happened", err) // same as above but we also get a date time stamp
		// fmt.Println("log.Fatalln()------------------------")
		// log.Fatalln("err happened", err) // Fatalln is equivalent to Println followed by a call to os.Exit(1).
		// fmt.Println("log.Panic() -------------------------")
		// log.Panic("err happened", err)
		fmt.Println("panic() -----------------------------")
		panic(err)  // program terminates after this, no recover functions are run
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
func foo() {
	fmt.Println("When os.Exit() is called, deferred functions dont't run.")
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.