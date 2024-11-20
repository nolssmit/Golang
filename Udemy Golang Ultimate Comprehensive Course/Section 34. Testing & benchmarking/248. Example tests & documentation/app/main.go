package main

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/acdc"
)

func main() {
	fmt.Println(acdc.Sum(9, 10))
	fmt.Println(acdc.Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

// example: https://pkg.go.dev/strings#ToUpper
// Also read: https://pkg.go.dev/testing#hdr-Examples
