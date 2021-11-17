package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// can be used for debugging, see the error flow/stack
		//stops normal execution of the current go routine
		panic(err)
	}
}
