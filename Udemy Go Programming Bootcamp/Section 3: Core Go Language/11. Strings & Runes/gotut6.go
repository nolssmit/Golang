// https://www.udemy.com/course/go-language/learn/lecture/35216204#notes
//
package main

import (
	//	"bufio"
	"fmt"
	"unicode/utf8"
	//	"log"
	//	"math/rand"
	//	"os"
	//	"strconv"
	"strings"
	// "time"
)

var pl = fmt.Println

func main() {
	// ----- STRINGS -----
	// Strings are arrays of bytes []byte
	// Escape Sequences : \n \t \" \\
	sV1 := "A word"

	// Replacer that can be used on multiple strings
	// to replace one string with another
	aReplacer := strings.NewReplacer("A", "Another")
	sV2 := aReplacer.Replace(sV1)
	pl(sV2)

	// Get length
	pl("Length :", len(sV2))

	// Contains string
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// Get first index match
	pl("o index :", strings.Index(sV2, "o"))

	// Replace all matches with 0
	// If -1 was 2 it would replace the 1st 2 matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	// Remove whitespace characters from beginning and end of string
	sV3 := "\nSome Words\n"
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)

	// Split at delimiter
	pl("Split :", strings.Split("a-b-c-d", "-"))

	// Upper and lowercase string
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	// Prefix or suffix
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// ----- RUNES -----
	// In Go characters are called Runes
	// Runes are unicodes that represent characters
	rStr := "abcdefg"

	// Runes in string
	pl("Rune Count :", utf8.RuneCountInString((rStr)))


	// Print runes in string
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}
