
// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 1) //buffered channel - with one space inside
	dataChan <- 789
	n := <- dataChan

	fmt.Printf("n = %d\n", n)
}
// No error because the sending and
// receiving may happen sequentially