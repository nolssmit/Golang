package main

import (
	"fmt"
	"math"
//	"math.Pi"
//	"math.Pow"
)

type square struct{
	length int
	width int
}

func (s square) calcArea() int {
	return s.length * s.width
}

type circle struct{
	radius float64
}

func (c circle) calcArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface{
	calcArea()
}

func calcSomeArea(s shape){
	s.calcArea()
}

func main(){
	sq := square{
		length: 6,
		width: 8,
	}
	ci := circle{
		radius: 100,
	}

	// fmt.Println("Area square =",sq.calcArea())
	// fmt.Printf("Area circle = %.2f\n", ci.calcArea())
	// fmt.Println("Area square =", calcSomeArea(sq))
	// fmt.Printf("Area circle = %.2f\n", calcSomeArea(ci))	

	calcSomeArea(sq)
}
