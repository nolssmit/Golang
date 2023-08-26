package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)
	numRounds := 3

	// Spawn a new thread, see: https://yourbasic.org/golang/goroutines-explained/
	go throwingNinjaStar(channel, numRounds)

	for i := 0; i < numRounds; i++ {
		fmt.Println(<-channel)
	}
}

func throwingNinjaStar(channel chan string, numRounds int) {
	
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numRounds; i++ {
		// Get random number     0 <= i < 10
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
}
