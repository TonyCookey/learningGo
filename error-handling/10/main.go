package main

import (
	"fmt"
	"log"
	"math"
)

type ErrNegNum struct {
	errType   string
	operation string
	err       error
}

func (e *ErrNegNum) Error() string {
	return fmt.Sprintln(e.errType, ":", e.operation, "==> ", e.err)
}

func main() {
	val, err := squareRoot(-36)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Square Root:", val)
}
func squareRoot(x float64) (float64, error) {
	if x <= 0 {
		return 0, &ErrNegNum{
			errType:   "Computation Error",
			operation: "Square Root",
			err:       fmt.Errorf("square root of a negative number (%v)", x),
		}
	}
	return math.Sqrt(x), nil
}
