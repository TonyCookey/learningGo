package main

import (
	"fmt"
	"learningGo/testing/01/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Convert(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.Convert(20))
}
