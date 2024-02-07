package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("Hello World!")
}

func TestAdd(*testing.T){
	total := Add(5,5)
	if total != 10 {
		t.Errorf("Sum was incorrect")
	}
}