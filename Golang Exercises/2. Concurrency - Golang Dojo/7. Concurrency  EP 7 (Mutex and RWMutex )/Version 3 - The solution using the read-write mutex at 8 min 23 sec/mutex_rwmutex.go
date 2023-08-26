package main

import (
	"fmt"
	"sync"
	"time"
)

// Package global variable
var (
	lock   sync.Mutex
	rwlock sync.RWMutex
	count  int
)

func main() {
	readAndWrite()
}

func increment() {
	lock.Lock()
	// Only one go routine has access to the global package variable count
	count++
	lock.Unlock()
}

func readAndWrite() {
	go read()
	go read()
	go write()
	go read()

	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}

func read() {
	rwlock.RLock()
	defer rwlock.RUnlock()

	fmt.Println("Read locking")
	time.Sleep(time.Second * 1)
	fmt.Println("Reading unlocking")
}

func write() {
	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write locking")
	time.Sleep(time.Second * 1)
	fmt.Println("Write unlocking")
}

