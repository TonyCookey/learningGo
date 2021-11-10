package main

import "fmt"

type Square struct {
	length float64
}
type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.length * s.length
}
func info(s Shape) float64 {
	return s.area()
}

func main() {
	sq1 := Square{length: 20}
	fmt.Println(info(sq1))
}
