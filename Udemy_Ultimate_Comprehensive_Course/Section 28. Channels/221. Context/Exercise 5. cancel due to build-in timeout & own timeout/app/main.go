// https://www.youtube.com/watch?v=VkGQFFl66X4&t=156s&ab_channel=CodeWithRyan
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {

	timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second *2)
	defer cancel()

	// create HTTP request
	reg, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://www.yahoo.com", nil)
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

	fmt.Printf("downloaded image of size %d,\n", len(imagedata))
	})
}
