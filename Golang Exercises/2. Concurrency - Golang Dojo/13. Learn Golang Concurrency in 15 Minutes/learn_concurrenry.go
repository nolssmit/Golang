package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// goroutine .. same as theads in other programming languages
	go func() {
		fmt.Println("New goroutine")
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("This is the old one")

	// unbuffered channel
	ch := make(chan string)
	go func() {
		ch <- "Hello World!"
	}()
	fmt.Println(<-ch)

	// buffered channel
	ch = make(chan string, 2)
	ch <- "1"
	ch <- "2"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// wait group
	var wg sync.WaitGroup
	wg.Add(1)   // the number of go routines we want to wait on the main go routine
	go func() { // spawn the the new routines we want to wait for in the main routine
		fmt.Println("New goroutine")
		wg.Done() // notify the following wait method so that the instructions ...
	}()
	wg.Wait() // ... following the wait routne will be called
	fmt.Println("This is the old one")

	// mutex (mutal exclusion on resources)
	iterations := 10000
	sum := 0
	var wg_mutex sync.WaitGroup
	wg_mutex.Add(iterations)
	var mu sync.Mutex                 // initialize mutex
	for i := 0; i < iterations; i++ { // for each iteration we spawn a new routine ...
		go func() { // ... and each routine increment the sum the number of iterations
			mu.Lock()       // lock mutex
			sum++           // the sum routine is not thread safe and must be protexted by a mutex
			mu.Unlock()     // unlock mutex ...
			wg_mutex.Done() // ... so that the main routine will be last route to finish executing
		}()
	}
	wg_mutex.Wait()
	fmt.Println("sum: ", sum)

	// once  ... to ensure one task is performed only once
	iterations_once := 100000
	sum_once := 0
	var wg_once sync.WaitGroup
	wg_once.Add(iterations_once)
	var once sync.Once                     // initialize mutex
	for i := 0; i < iterations_once; i++ { // for each iteration we spawn a new routine ...
		go func() { // ... and each routine increment the sum once
			once.Do(func() {
				sum_once++
			})
		}()
		wg_once.Done() // so that the main routine will be last route to finish executing
	}
	wg_once.Wait()
	fmt.Println("sum_once: ", sum_once)

	// pool            ...   we utilize resource we previously allocated
	memPool := &sync.Pool{
		New: func() interface{} { // new member method returning a interface type
			mem := make([]byte, 1024) // we allocate one MB of memory
			return &mem               // we return that piece of memory
		},
	}
	mem := memPool.Get().(*[]byte) // we return the first piece of memory we allocated and cast it into the type we want
	// ... some instructions to utilize this piece of memory
	// ...
	memPool.Put(mem) // put the memory back into the memory pool and next time we call the Get method to utilize the allocated memory

	//cond     .... To signal or broadcat between go routines
	c := sync.NewCond(&sync.Mutex{})

	go func() {
		c.L.Lock()
		// instruction to change some condition
		c.L.Unlock()
		c.Signal() // or c.Broadcast()
	}()
	// main routine checking condition
	// if .. {

	// }
	c.L.Lock()
	c.Wait()
	c.L.Unlock()



	// map
	const iterations_map = 1000
	var wg_map sync.WaitGroup   // a thread-safe or go-routine safe map
	syncMap := sync.Map{}       // calling the sync map library
	wg_map.Add(iterations_map)  // adding a thousand iterations to the wait group

	for i := 0; i < iterations_map; i++ {  // spwning a go routine for each iteration
		go func() {
			syncMap.Store(0, i)  // with the parameters the key and value
			wg_map.Done()
		}()
	}
	wg_map.Wait()   // waiting for a thousand routines to complete



	var i int64
	atomic.AddInt64(&i, 1)
	// equivalent of
	// mu.Lock()
	// i =+ 1
	// mu.Unlock()
	var av atomic.Value  // perform atomic operations on custom values
	type ninja struct {   // a custom type
		name string       // member variable
	}
	av.Store(ninja{"Wallece"})  // atomicly store different ninja objects
	av.Load()                   // getting the stored value

}
