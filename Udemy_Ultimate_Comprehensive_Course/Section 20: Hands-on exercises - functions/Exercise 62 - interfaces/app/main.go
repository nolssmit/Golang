package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s1 := square{
		length: 6,
		width:  8,
	}

	c1 := circle{
		radius: 100,
	}

	fmt.Println("Area square =", s1.area())
	fmt.Printf("Area circle = %.2f\n", c1.area())

	fmt.Println("Area of rectangle is:", info(s1))
	fmt.Printf("Area of circle is: %.2f\n", info(c1))
}
