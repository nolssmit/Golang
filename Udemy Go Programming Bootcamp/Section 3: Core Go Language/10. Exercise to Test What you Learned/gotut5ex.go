// https://www.udemy.com/course/go-language/learn/lecture/35216192#notes
//
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)

	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random number generated: ", randNum)

	fmt.Print("Guess a random number between 0 and 50 :")

	for true {
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Trims whitespace from guess
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if randNum == iGuess {
			fmt.Println("You guessed it!")
			break
		} else if iGuess > randNum {
			fmt.Println("Guess something lower!")
		} else {
			fmt.Println("Guess something higher!")
		}
	}
}
