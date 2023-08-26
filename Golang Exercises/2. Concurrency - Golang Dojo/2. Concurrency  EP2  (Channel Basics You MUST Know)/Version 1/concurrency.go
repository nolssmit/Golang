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

	evilNinja := "Tommy"
	go attack(evilNinja)

	time.Sleep(time.Second * 2)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Thowing ninja stars at ", target)
}
