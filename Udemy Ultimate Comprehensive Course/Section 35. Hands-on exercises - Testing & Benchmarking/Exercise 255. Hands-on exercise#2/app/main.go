package main

import (
	"fmt"
	"github.com/nolssmit/Golang/GolangPackages/quote"
	"github.com/nolssmit/Golang/GolangPackages/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}