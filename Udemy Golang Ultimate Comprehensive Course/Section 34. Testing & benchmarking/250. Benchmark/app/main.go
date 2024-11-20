package main

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/saying"
)

func main() {
	
	fmt.Println(saying.Greet("Nols Smit"))
}

// From package sayings's directory: $ go test
// To view documentation: $ godoc -http=:8080
// ... and in browser: http://localhost:8080
// From package sayings's directory, 
// ... do a benchmark on all functions: $ go test -bench=.
// ... do a benchmark on Greet(): $ go test -bench Greet