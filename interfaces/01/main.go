package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.length * s.length
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(s Shape) float64 {
	return s.area()
}

func main() {
	sq1 := Square{length: 20}
	c1 := Circle{radius: 30}
	fmt.Println("Square:", info(sq1))
	fmt.Println("Circle:", info(c1))
}
