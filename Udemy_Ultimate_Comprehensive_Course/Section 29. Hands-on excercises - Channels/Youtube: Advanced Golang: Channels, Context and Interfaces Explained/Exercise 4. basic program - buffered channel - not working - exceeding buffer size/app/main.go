
// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 1) //buffered channel - with one space inside
	dataChan <- 789  // add data to channel
	dataChan <- 790  // add data to channel

	n := <- dataChan
	fmt.Printf("n = %d\n", n)
}
// error because then buffer is one
// and you can't add more then one