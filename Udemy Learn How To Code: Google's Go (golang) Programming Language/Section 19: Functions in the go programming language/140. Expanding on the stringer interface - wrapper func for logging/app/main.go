package main

import (
	"fmt"
	"log"
	"strconv"
)

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM 138", s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	// fmt.Println(b)
	// fmt.Println(n)
	// log.Println(b)
	// log.Println(n)
	logInfo(b)
	logInfo(n)
}
