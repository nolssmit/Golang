package main

import (
	"fmt"
	"strconv"
)

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

type book struct{
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book {
		title: "West With The Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)
}
// https://pkg.go.dev/fmt#Stringer