// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	dataChan := make(chan int) // unbuffered channel - unlimited capacity
	go func() {
		//defer close(dataChan) // implys we are done with the channel
		wg := sync.WaitGroup{}  // to track when the go routines are done
		for i := 0; i < 1000; i++ {
			wg.Add(1)  // for each iteration of the loop
			go func() {  // create a new go routine for each iteration of the loop
				defer wg.Done()  // will execute upon completion of the go routine
				result := doWork()
				dataChan <- result
			}()
		}
		wg.Wait()  // wait till all the iterations are done before closing the channel
		close(dataChan)  // implys we are done with the channel
	}()

	//simultaneously read from the channel 
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}
}

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

// using an anonymous function to call the doWork() function for each iteration
// working, because anonymous function excited and closing the channel
// therefore telling the main function when to stop the loop
