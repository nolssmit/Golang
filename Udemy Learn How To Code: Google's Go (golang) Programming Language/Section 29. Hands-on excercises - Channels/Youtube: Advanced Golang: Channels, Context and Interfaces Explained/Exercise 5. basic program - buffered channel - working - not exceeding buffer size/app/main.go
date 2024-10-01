
// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 2) //buffered channel - with two spaces inside
	dataChan <- 789  // add data to channel
	dataChan <- 790  // add data to channel

	n := <- dataChan
	fmt.Printf("n = %d\n", n)

	n = <- dataChan
	fmt.Printf("n = %d\n", n)
}
// no error