package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println((time.Since(now)))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	// Note: This defined smokeSignal channel can be used by only one ninja
	// and the following statement will cause a deadlock error and can be be relved by using buffered channels
	// smokeSignal <- false
	go attack(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Thowing ninja stars at ", target)
	attacked <- true
}
