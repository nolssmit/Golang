package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)	// Rename the file to see the error
		return
	}

	fmt.Println(string(bs))
}