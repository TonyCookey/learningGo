package main

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
	fmt.Println(s.area())
}

func main() {
	c1 := circle{radius: 12}
	info(c1)
	//a value receiver will also accept a pointer type
	info(&c1)
}
