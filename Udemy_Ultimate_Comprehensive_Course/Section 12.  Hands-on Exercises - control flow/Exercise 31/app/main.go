package main

import (
	"fmt"
	"time"
)

func main() {
	i := 100000000
	startTime := time.Now()
	for  i >= 0  {
		i--	
	}
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Time processed: %v\n", executionTime)
}