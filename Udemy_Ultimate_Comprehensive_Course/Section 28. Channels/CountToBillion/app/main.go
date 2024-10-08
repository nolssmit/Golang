package main

import (
	"fmt"
	"time"
)


var j = 0
var start = time.Now()
func main() {
	for i:=0; i<1000000000; i++ {
		j++
	}
    timeLapsed := time.Since(start)
	fmt.Println("Runtime: ",timeLapsed.Round(time.Second))
}