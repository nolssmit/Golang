package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// Write interface: https://pkg.go.dev/io#Writer

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("err %s", err)
	}
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}

// see output.txt
