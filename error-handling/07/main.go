package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

//ErrNegNum - idiomatic way of define error and starting the error with Capital letters and Err
var ErrNegNum = errors.New("computation error: square root of a negative number")

func main() {
	val, err := squareRoot(-36)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Square Root:", val)
}

func squareRoot(x float64) (float64, error) {
	if x <= 0 {
		return 0, ErrNegNum
	}
	return math.Sqrt(x), nil
}
