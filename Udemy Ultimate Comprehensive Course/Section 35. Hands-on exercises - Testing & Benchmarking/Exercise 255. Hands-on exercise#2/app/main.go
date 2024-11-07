package main

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/word"
	"github.com/nolssmit/Golang/GolangPackages/quote"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}