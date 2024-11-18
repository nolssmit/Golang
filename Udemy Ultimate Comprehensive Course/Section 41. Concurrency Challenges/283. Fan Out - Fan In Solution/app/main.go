package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	input := make(chan string)
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1* time.Millisecond)
}

// publisher pushes data into a channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data \n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}
	
func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input ...\n", thisID)
		input := <- in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}
// Install: $ go get -u golang.org/x/lint/golint

/*
CHALLENGE #1
Is this fan out?
-- YES
Are we "fanning out" work? Yes. We've launched several goroutines that are simultanously publishing a
message anto our channel. The golang blog says, "Fan out means you have multiple functions reading from
the same channel until that channel is closed." Here we do have multiple functions reading from the
same channel. So, okay, we're fanning out.

CHALLENGE #2
Is this fan in?
-- NO
What is being "fanned in" here? We have lauched several goroutines of the same function: workerProcess.
What do those goroutines do? They are all reading from an unbuffered channel. If there was a tremendous
amount of processing that each "workerProcess" func executed, then all three of the "workerProcess" 
funcs could be processing in parallel: pulling values off the channel and processing them. There is no
"fanning in" though here. Remember what the golang blog describes fan in: "A function can read from
multiple inputs and proceed until all are closed by multiplexing the input channels into a single one channel."
*/

