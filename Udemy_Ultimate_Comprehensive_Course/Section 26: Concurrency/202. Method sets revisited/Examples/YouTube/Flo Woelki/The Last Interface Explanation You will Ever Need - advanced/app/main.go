// https://www.youtube.com/watch?v=SX1gT5A9H-U&t=569s
package main

import (
	"fmt"
	"math"
)

// an interface is a contract that a type must satisfy in order to be used in the contract
type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

// Interface composition - embed one interface into another
type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func describeShape(g Geometry) {
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Perimeter())
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}

func main() {
	rect := Rectangle{width: 10, height: 10}
	describeShape(rect)

}

type CalculationError struct {	
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{msg: "Invalid input",}
	}
	return math.Sqrt(val), nil
}