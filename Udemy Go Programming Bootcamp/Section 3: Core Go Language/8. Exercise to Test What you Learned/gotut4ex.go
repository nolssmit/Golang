// https://www.udemy.com/course/go-language/learn/lecture/35216180#notes

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Create alias to long function names
var pl = fmt.Println

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	//-------------- Read number 1 -------------
	fmt.Print("Enter Number 1 : ")

	num1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trims whitespace from guess
	num1 = strings.TrimSpace(num1)
	iNum1, err := strconv.Atoi(num1)

	if err != nil {
		log.Fatal(err)
	}
	//-------------- Read number 2 -------------
	fmt.Print("Enter Number 2 : ")

	// Trims whitespace from guess
	num2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	num2 = strings.TrimSpace(num2)
	iNum2, err := strconv.Atoi(num2)

	if err != nil {
		log.Fatal(err)
	}

	// Use Printf to create formatted string
	sum := iNum1 + iNum2
	fmt.Printf("%d + %d = %d\n", iNum1, iNum2, sum)
}
