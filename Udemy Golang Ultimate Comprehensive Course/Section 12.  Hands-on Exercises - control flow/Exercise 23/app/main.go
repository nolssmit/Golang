package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("../SNOWY-EVENING.txt")
	if err != nil {
		fmt.Print(err)
	}

	h := sha256.New()
	h.Write([]byte(b))
	fmt.Printf("%x", h.Sum(nil))
}
