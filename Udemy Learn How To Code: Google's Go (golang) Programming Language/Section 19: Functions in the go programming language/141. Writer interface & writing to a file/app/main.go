package main

import (
	"log"
	"os"
)

// Write interface: https://pkg.go.dev/io#Writer

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("err %s", err)
	}
	defer f.Close()

	s := []byte("Hello gophers!")

	_, err = f.Write(s) //err already decleared, now re-assigning
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

// see output.txt