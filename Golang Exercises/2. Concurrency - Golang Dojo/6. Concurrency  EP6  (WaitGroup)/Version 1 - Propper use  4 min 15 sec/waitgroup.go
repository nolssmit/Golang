package main

import (
	"fmt"
	"sync"
)

func main() {
	var beeper sync.WaitGroup
	// Slice of strings
	var evilNinjas = []string {"Tommy", "Johnny", "Bobby"}
	beeper.Add(len(evilNinjas))
	// Note we pass the address of beeper
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}
	
	beeper.Wait()
	fmt.Println("Mission completed")
}

// Note the pointer of type sync.WaitGroup
func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}