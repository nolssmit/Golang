package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool

func main() {
//	gettingReadyForMission()
	gettingReadyForMissionWithCond()
}

func gettingReadyForMission() {
	go gettingReady()
	workingIntervals := 0

	for !ready {
		time.Sleep(5 * time.Second)
		workingIntervals++
	}
	fmt.Printf("We are now ready! After %d work intervals.\n", workingIntervals)
}

func gettingReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)
	workingIntervals := 0

	cond.L.Lock()
	for !ready {
		workingIntervals++
		cond.Wait()
	}

	cond.L.Unlock()

	fmt.Printf("We are now ready! After %d work intervals.\n", workingIntervals)
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep() // taking some time to get ready
	ready = true
	cond.Signal()
}

func gettingReady() {
	sleep() // taking some time to get ready
	ready = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1 + rand.Intn(5)) * time.Second
	time.Sleep((someTime))
}
