// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// a timeout less than 2 seconds will result in a panic
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*2000)
	defer cancel()

	// create HTTP request
	reg, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/2000x2000", nil)
	if err != nil {
		panic(err)
	}

	// perform HTTP request
	resp, err := http.DefaultClient.Do(reg)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// get data from HTTP response
	imagedata, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("downloaded image of size %d\n", len(imagedata))
}
