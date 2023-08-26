package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

const iterations = 1000

func main() {
	var sum int64
	fmt.Println(sum)
	atomic.AddInt64(&sum, 1)
	fmt.Println(sum)

	var mu sync.Mutex
	mu.Lock()
	sum = sum + 1
	mu.Unlock()
	fmt.Println(sum)

	var diffSum int64
	atomic.LoadInt64(&diffSum)
	fmt.Println(atomic.LoadInt64(&diffSum))
	atomic.StoreInt64(&diffSum, 1)
	fmt.Println(diffSum)

	var av atomic.Value
	wallace := ninja{"Wallece"}
	av.Store((wallace))

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		w := av.Load().(ninja)
		w.name = "Not Wallace"
		av.Store(w)
		w.name = "Wallace Again"
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(av.Load().(ninja).name)
}

type ninja struct{
	name string
}

func BenchmarkMutexAdd(t *testing.B) {
	var wg sync.WaitGroup
	wg.Add(iterations)
	var sum int64
	var mutex sync.Mutex
	for i := 0; i < iterations; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				mutex.Lock()
				sum++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicAdd(t *testing.B) {
	var wg sync.WaitGroup
	wg.Add(iterations)
	var sum int64
	for i := 0; i < iterations;  i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				atomic.AddInt64(&sum, 1)
			}
		}()
	}
	wg.Wait()
}