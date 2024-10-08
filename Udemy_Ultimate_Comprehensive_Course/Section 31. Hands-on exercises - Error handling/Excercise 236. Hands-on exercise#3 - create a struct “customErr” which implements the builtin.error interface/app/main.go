package main

import (
	"fmt"
	"log"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here's the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Need more coffee",
	}
	foo(c1)
}

func foo(e error) {
	if e != nil {
		log.Println("foo ran -", e, "\n", e.(customErr).info)
	}
}
// Type assertions in Structs
// https://medium.com/@jamal.kaksouri/mastering-type-assertion-in-go-a-comprehensive-guide-216864b4ea4d
// https://go.dev/tour/methods/15