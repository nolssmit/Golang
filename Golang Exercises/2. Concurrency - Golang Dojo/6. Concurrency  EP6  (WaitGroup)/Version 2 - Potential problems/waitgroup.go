package main

import (
	"sync"
)

func main() {
	var beeper sync.WaitGroup

	beeper.Add(1)
	// Deadlock without beeper.Done()
	beeper.Done()
	// Negeative WaitGroup counter with extra beeper.Done()
	// beeper.Done()
	beeper.Wait()
}
