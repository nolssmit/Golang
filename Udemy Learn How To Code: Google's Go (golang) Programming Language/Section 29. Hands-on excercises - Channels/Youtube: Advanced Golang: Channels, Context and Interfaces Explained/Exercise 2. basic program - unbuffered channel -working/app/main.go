
// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int) // unbuffered channel - no space inside
	go func(){
		dataChan <- 789
	}()
	
	n := <- dataChan

	fmt.Printf("n = %d\n", n)
}
// working because anonymous function and main function 
// execute simultaneously in two different threads
