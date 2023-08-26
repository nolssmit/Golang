package main

import (
	"fmt"
	"sync"
	"testing"
)

const iterations = 1000

func BenchmarkRegularMap(t *testing.B) {
	var mutex sync.Mutex
	regularMap := make(map[int]int, iterations)
	var wg sync.WaitGroup
	wg.Add(iterations - 1)
	regularMap[0] = 0

	for i := 1; i < iterations; i++ {
		go func() {
			defer wg.Done()

			for j := 1; j < iterations; j++ {
				go func() {
					mutex.Lock()
					defer mutex.Unlock()
					regularMap[i] = j
				}()
			}
		}()
	}
	wg.Wait()
}

func BenchmarkSyncMap(t *testing.B) {
	var syncMap sync.Map 
	var wg sync.WaitGroup
	wg.Add(iterations -1)
	syncMap.Store(0, 0)

	for i := 1; i < iterations; i++ {
		go func() {
			defer wg.Done()

			for j := 1; j < iterations; j++ {
				go func() {
					syncMap.Store(i, j)
				}()
			}
		}()
	}

	wg.Wait()
}

func main() {
	regularMap := make(map[int]interface{})
	syncMap := sync.Map{}

	// put
	regularMap[0] = 0
	regularMap[1] = 1
	regularMap[2] = 2
	syncMap.Store(0, 0)
	syncMap.Store(1, 1)
	syncMap.Store(2, 2)

	// get
	regularValue, regularOk := regularMap[0]
	fmt.Println(regularValue, regularOk)
	syncValue, syncOk := syncMap.Load(0)
	fmt.Println(syncValue, syncOk)

	// delete
	regularMap[1] = nil
	syncMap.Delete(1)
	
	// get and delete
	syncValue, loaded := syncMap.LoadAndDelete(2)
	mu := sync.Mutex{}
	mu.Lock()
	regularValue = regularMap[2]
	delete(regularMap, 2)
	mu.Unlock()
	fmt.Println(syncValue, loaded, regularValue)

	// get and put
	syncValue, loaded = syncMap.LoadOrStore(1, 1)
	mu = sync.Mutex{}
	mu.Lock()
	regularValue, regularOk = regularMap[1]
	if regularOk {
		regularMap[1] = 1
		regularValue = regularMap[1]
	}
	mu.Unlock()
	fmt.Println(syncValue, regularValue)

	// range
	for key, value := range regularMap {
		fmt.Print(key, value, " | ")
	}

	fmt.Println()
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Print(key, value, " | ")
		return true
	})	
}
