package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ready bool

// Note: You are broadcating to muliple go routines
func main() {
	gettingReadyForMissionWithCond()
	broadcastStartOfMission()
}

func broadcastStartOfMission() {
	beeper := sync.NewCond((&sync.Mutex{}))
	var wg sync.WaitGroup
	wg.Add(3)
	standByForMission(func() {
		fmt.Println("Ninja 1 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 2 starting mission.")
		wg.Done()
	}, beeper)
	standByForMission(func() {
		fmt.Println("Ninja 3 starting mission.")
		wg.Done()
	}, beeper)
	beeper.Broadcast()
	wg.Wait()
	fmt.Println("All Ninjas have started their missions")
}

func standByForMission(fn func(), beeper *sync.Cond) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		beeper.L.Lock()
		defer beeper.L.Unlock()
		beeper.Wait()
		fn()
	}()
	wg.Wait()
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
