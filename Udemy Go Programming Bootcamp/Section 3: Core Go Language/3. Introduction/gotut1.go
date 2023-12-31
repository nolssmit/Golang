// https://www.udemy.com/course/go-language/learn/lecture/35216118#questions/19276818

// A package is a collection of code
// We can define what package we want our code to belong to
// We use main when we want our code to run in the terminal
package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Create alias to long function names
var pl = fmt.Println

// When a Go program executes it executes a function named main
// Go statements don't require semicolons
func main() {
	// Prints text and a newline
	// List package name followed by a period and the function name	
	pl("Hello Go")

	// Get user input (To run this in the terminal go run hellogo.go)	
	fmt.Print("What is your name? ")

	// Setup buffered reader that gets text from the keyboard (os.Stdin)	
	reader := bufio.NewReader(os.Stdin)
	
	// Copy text up to the newline
	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')
	// It is better to handle it	
	name,err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)	
	} else {
		log.Fatal(err)
	}

}
