// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int) // unbuffered channel - unlimited capacity
	go func(){
		defer close(dataChan)	// implys we are done with the channel
		for i := 0; i < 10; i++ {
			dataChan <- i
		}
	}()
	
	//simultaneously read from the channel
	for n := range dataChan {
		fmt.Printf("n = %d\n", n)
	}

}
// working, because anonymous function excited and closing the channel
// therefore telling the main function when to stop the loop