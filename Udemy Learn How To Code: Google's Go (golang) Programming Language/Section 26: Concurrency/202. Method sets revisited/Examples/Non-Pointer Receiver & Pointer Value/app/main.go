package main
// https://go.dev/play/p/glWZmm0gY6
// Non-Pointer Receiver & Pointer Value

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
  	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}

