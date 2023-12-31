package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

// ------- CONCURRENCY --------
func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}

// These functions will print in order using channels
// Func receives a channel and then sends values
// over channels once each time it is called
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// ----- BANK ACCOUNT EXAMPLE -----
// Here I'll simulate customers accessing a bank account 
// and lock out customers to allow for individual access
type Account struct {
	balance int
	lock 	sync.Mutex // Mutal exclusion: Lock out other users when accessing
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock() //Unlock when it's able to for other users to access it
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock() //Unlock when it's able to
	if v > a.balance {
		pl("Not Enough Money in Account")
	} else {
		a.balance -= v
		fmt.Printf("%d withdrawn : Balance : %d\n",
				v, a.balance)
		//a.balance -= v
	}
}

func main() {	
	// ----- CONCURRENCY -----
	// Concurrency allows us to have multiple blocks of code share execution 
	// time by pausing their execution. We can also run blocks of codes in 
	// parallel at the same time. 
	// (Concurrent tasks in Go are called goroutines)

	// To execute multiple functions in new goroutines add the word go in 
	// front of the function calls (Those functions can't have return values)

	// We can't control when functions execute so we may get different results
	go printTo15()
	go printTo10()

	// We have to pause the main function because if main ends so 
	// will the goroutines
	time.Sleep(2 * time.Second) // Pause 2 seconds

	// You can have goroutines communicate by using channels. 
	// The sending goroutine also makes sure the receiving goroutine
	// receives the value before it attempts to use it

	// Create a channel : Only carries values of 1 type
	channel1 := make(chan int)
	channel2 := make(chan int)

	pl("---Calling the channels---")
	go nums1(channel1)
	go nums2(channel2)

	pl("---Execute channel1 3x---")
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)

	pl("---Execute channel2 3x---")
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)

	pl("---change the order end execute---")
	go nums1(channel1)
	go nums2(channel2)	
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel1)	
	pl(<-channel2)
	pl(<-channel1)
	pl(<-channel2)


pl("---Simulate a bank account operations---")
	// Using locks to protect data from being accessed by more than one user 
	// at a time 
	// Locks are another option when you don't have to pass data
	var acct Account
	acct.balance = 100
	pl("Starting Balance :",acct.GetBalance())
	// we create a bunch of go routines to try to withdraw money from 
	// bank account
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}