// https://www.youtube.com/watch?v=SX1gT5A9H-U&t=569s
package main

import (
	"fmt"
)

// See a interface as the contract
type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}

func main() {
	rect := Rectangle{width: 10, height: 10}
	circle := Circle{Radius: 10}
	fmt.Println("Area of rectangle: ", calculateArea(rect))
	fmt.Println("Area of circle: ", calculateArea(circle))

	// Mystery box holds the interface value of 10
	mysterBox := interface{}(10)
	describeValue(mysterBox)

	retrievedInt, ok := mysterBox.(int)
	if ok {
		fmt.Println("retrievedInt: ", retrievedInt)
	} else {
		fmt.Println("mysterBox is not an int")
	}
}
