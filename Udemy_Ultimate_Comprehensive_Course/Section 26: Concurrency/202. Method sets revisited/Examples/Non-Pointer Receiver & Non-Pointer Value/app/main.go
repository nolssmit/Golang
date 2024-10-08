package main
// https://go.dev/play/p/2ZU0QX12a8
// Non-Pointer Receiver & Non-Pointer Value

import (
	"fmt"
	"math"
)

type circle struct{
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info (s shape) {
	fmt.Println("Area:", s.area())
}

func main(){
  c := circle{5}
  info(c)
}